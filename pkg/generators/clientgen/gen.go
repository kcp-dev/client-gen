/*
Copyright 2022 The KCP Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package clientgen

import (
	"bytes"
	"errors"
	"fmt"
	"go/format"
	"io"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"
	"k8s.io/code-generator/cmd/client-gen/args"
	"k8s.io/code-generator/cmd/client-gen/types"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/loader"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/pkg/flag"
	"github.com/kcp-dev/code-generator/pkg/internal/clientgen"
	"github.com/kcp-dev/code-generator/pkg/util"
	genutil "k8s.io/code-generator/cmd/client-gen/generators/util"
)

const (
	// GeneratorName is the name of the generator.
	GeneratorName = "client"
	// packageName for typed client wrappers.
	typedPackageName = "typed"
	// name of the file while wrapped clientset is written.
	clientSetFilename = "clientset.go"
)

var (
	// In controller-tool's terms marker's are defined in the following format: <makername>:<parameter>=<values>. These
	// markers are not a part of genclient, since they do not accept any values.
	nonNamespacedMarker = markers.Must(markers.MakeDefinition("genclient:nonNamespaced", markers.DescribesType, struct{}{}))
	noStatusMarker      = markers.Must(markers.MakeDefinition("genclient:noStatus", markers.DescribesType, struct{}{}))
	noVerbsMarker       = markers.Must(markers.MakeDefinition("genclient:noVerbs", markers.DescribesType, struct{}{}))
	readOnlyMarker      = markers.Must(markers.MakeDefinition("genclient:readonly", markers.DescribesType, struct{}{}))

	// These markers, are not a part of "+genclient", and are defined separately because they accept a list which is comma separated. In
	// controller-tools, comma indicates another argument, as multiple arguments need to provided with a semi-colon separator.
	skipVerbsMarker = markers.Must(markers.MakeDefinition("genclient:skipVerbs", markers.DescribesType, markers.RawArguments("")))
	onlyVerbsMarker = markers.Must(markers.MakeDefinition("genclient:onlyVerbs", markers.DescribesType, markers.RawArguments("")))
)

type genclient struct {
	Method      *string
	Verb        *string
	Subresource *string
	Input       *string
	Result      *string
}

type Generator struct {
	// inputDir is the path where types are defined.
	inputDir string
	// inputpkgPaths stores details on input directory.
	inputpkgPaths pkgPaths
	// outputpkgPaths stores details on output directory.
	outputpkgPaths pkgPaths
	// output Dir where the wrappers are to be written.
	outputDir string
	// path to where generated clientsets are found.
	clientSetAPIPath string
	// path to optional apply-config-gen packages.
	applyConfigGenPkg string
	// clientsetName is the name of the generated clientset package.
	clientsetName string
	// GroupVersions for whom the clients are to be generated.
	groupVersions []types.GroupVersions
	// headerText is the header text to be added to generated wrappers.
	// It is obtained from `--go-header-text` flag.
	headerText string
	// clusterclientWrappers contains the content which has the
	// initial config and wrappers.
	clusterclientWrappers []byte
}

// TODO: Store this information in generation context, as other genrators
// may need this too.
type pkgPaths struct {
	// basePacakge path as found in go module.
	basePackage string
	// hasGoMod is a way of checking if the go.mod file is present inside
	// the input directory or not. If present the basepkg path need not be modified
	// to include the location of input directory. If not, include the location of
	// all the sub folders provided in the input directory.
	hasGoMod bool
}

func (g Generator) RegisterMarker() (*markers.Registry, error) {
	reg := &markers.Registry{}
	if err := markers.RegisterAll(
		reg,
		GenclientMarker,
		nonNamespacedMarker,
		noStatusMarker,
		noVerbsMarker,
		readOnlyMarker,
		skipVerbsMarker,
		onlyVerbsMarker,
	); err != nil {
		return nil, fmt.Errorf("error registering markers")
	}
	return reg, nil
}

func (g Generator) GetName() string {
	return GeneratorName
}

// Run validates the input from the flags and sets default values, after which
// it calls the custom client genrator to create wrappers. If there are any
// errors while generating interface wrappers, it prints it out.
func (g Generator) Run(ctx *genall.GenerationContext, f flag.Flags) error {
	if err := flag.ValidateFlags(f); err != nil {
		return err
	}
	if err := g.setDefaults(f); err != nil {
		return err
	}
	if err := g.generate(ctx); err != nil {
		return err
	}

	// print all the errors consolidated from packages in the generation context.
	// skip the type errors since they occur when input path does not contain
	// go.mod files.
	hadErr := loader.PrintErrors(ctx.Roots, packages.TypeError)
	if hadErr {
		return fmt.Errorf("generator did not run successfully")
	}
	return nil
}

// setDefaults sets the default values for the generator. It also creates
// a list of group versions provided as an input.
func (g *Generator) setDefaults(f flag.Flags) (err error) {
	if f.InputDir != "" {
		g.inputDir = f.InputDir
		pkg, hasGoMod := util.CurrentPackage(f.InputDir)
		if len(pkg) == 0 {
			return fmt.Errorf("error finding the module path for this package %q", f.InputDir)
		}
		g.inputpkgPaths = pkgPaths{
			basePackage: pkg,
			hasGoMod:    hasGoMod,
		}
	}
	if f.OutputDir != "" {
		pkg, hasGoMod := util.CurrentPackage(f.OutputDir)
		if len(pkg) == 0 {
			return fmt.Errorf("error finding the module path for this package %q", f.OutputDir)
		}
		g.outputpkgPaths = pkgPaths{
			basePackage: pkg,
			hasGoMod:    hasGoMod,
		}
		g.outputDir = f.OutputDir
	}
	if f.ClientsetAPIPath != "" {
		g.clientSetAPIPath = f.ClientsetAPIPath
	}
	if f.ClientsetName != "" {
		g.clientsetName = f.ClientsetName
	}
	if f.ApplyConfigurationPackage != "" {
		g.applyConfigGenPkg = f.ApplyConfigurationPackage
	}
	g.headerText, err = util.GetHeaderText(f.GoHeaderFilePath)
	if err != nil {
		return err
	}
	gvs, err := GetGV(f)
	if err != nil {
		return err
	}
	g.groupVersions = append(g.groupVersions, gvs...)
	return nil
}

// GetGV parses the Group Versions provided in the input through flags
// and creates a list of []types.GroupVersions.
func GetGV(f flag.Flags) ([]types.GroupVersions, error) {
	groupVersions := make([]types.GroupVersions, 0)
	// Its already validated that list of group versions cannot be empty.
	gvs := f.GroupVersions
	for _, gv := range gvs {
		// arr[0] -> group, arr[1] -> versions
		arr := strings.Split(gv, ":")
		if len(arr) != 2 {
			return nil, fmt.Errorf("input to --group-version must be in <group>:<versions> format, ex: rbac:v1. Got %q", gv)
		}

		versions := strings.Split(arr[1], ",")
		for _, v := range versions {
			// input path is converted to <inputDir>/<group>/<version>.
			// example for input directory of "k8s.io/client-go/kubernetes/pkg/apis/", it would
			// be converted to "k8s.io/client-go/kubernetes/pkg/apis/rbac/v1".
			input := filepath.Join(f.InputDir, arr[0], v)
			groups := []types.GroupVersions{}
			builder := args.NewGroupVersionsBuilder(&groups)
			_ = args.NewGVPackagesValue(builder, []string{input})

			groupVersions = append(groupVersions, groups...)

		}
	}
	return groupVersions, nil
}

// generate first generates the wrapper for all the interfaces provided in the input.
// Then for each type defined in the input, it recursively wraps the subsequent
// interfaces to be kcp-aware.
func (g *Generator) generate(ctx *genall.GenerationContext) error {
	if err := g.writeWrappedClientSet(); err != nil {
		return err
	}
	return g.generateSubInterfaces(ctx)
}

func (g *Generator) writeWrappedClientSet() error {
	var out bytes.Buffer
	if err := g.writeHeader(&out); err != nil {
		return err
	}

	// Get the location of the typed wrapped clientset for imports.
	// Cases handled here, for example the scenarios could be:
	// Case 1:
	// if basePkg := k8s.io/kcp-dev; outputPkg := k8s.io/kcp-dev/output/examples
	// then typedPkgPath is k8s.io/kcp-dev/output/examples/
	// Case 2:
	// if basePkg := k8s.io/kcp-dev; outputPkg := ./output/examples
	// then typedPkgPath is k8s.io/kcp-dev/output/examples/
	// Case 3:
	// if basePkg := k8s.io/kcp-dev; outputPkg := .
	// then typedPkgPath is k8s.io/kcp-dev
	var typedPkgPath string
	if !g.outputpkgPaths.hasGoMod {
		typedPkgPath = filepath.Join(util.GetCleanRealtivePath(g.outputpkgPaths.basePackage, filepath.Clean(g.outputDir)), g.clientsetName)
	} else {
		typedPkgPath = filepath.Join(g.outputpkgPaths.basePackage, g.clientsetName)
	}

	wrappedInf, err := clientgen.NewInterfaceWrapper(g.clientSetAPIPath, g.clientsetName, typedPkgPath, g.groupVersions, &out)
	if err != nil {
		return err
	}

	if err := wrappedInf.WriteContent(); err != nil {
		return err
	}
	outBytes := out.Bytes()
	formattedBytes, err := format.Source(outBytes)
	if err != nil {
		return err
	} else {
		outBytes = formattedBytes
	}

	g.clusterclientWrappers = outBytes
	return nil
}

func (g *Generator) writeHeader(out io.Writer) error {
	n, err := out.Write([]byte(g.headerText))
	if err != nil {
		return err
	}

	if n < len([]byte(g.headerText)) {
		return errors.New("header text was not written properly.")
	}
	return nil
}

func (g *Generator) generateSubInterfaces(ctx *genall.GenerationContext) error {
	for _, gv := range g.groupVersions {
		// Each types.GroupVersions will have only one version.
		// Even if there are multiple versions for same group, we will have separate types.GroupVersions
		// for it. Hence length of gv.Versions will always be one.
		version := gv.Versions[0]

		// This is to accomodate the usecase wherein the apis are defined under a sub-folder inside
		// base package.
		basePkg := g.inputpkgPaths.basePackage
		if !g.inputpkgPaths.hasGoMod {
			cleanPkgPath := util.CleanInputDir(g.inputDir)
			if cleanPkgPath != "" {
				basePkg = filepath.Join(g.inputpkgPaths.basePackage, cleanPkgPath)
			}
		}

		path := filepath.Join(basePkg, gv.Group.String(), string(version.Version))

		pkgs, err := loader.LoadRootsWithConfig(&packages.Config{Dir: g.inputDir}, path)
		if err != nil {
			return err
		}

		// Assign the pkgs obtained from loading roots to generation context.
		// TODO: Figure out if controller-tools generation runtime can be used to
		// wire in instead.
		ctx.Roots = pkgs

		for _, root := range pkgs {
			root.NeedTypesInfo()

			// this is to accomodate multiple types defined in single group
			byType := make(map[string][]byte)

			importList := make([]string, 0)

			var outCommonContent bytes.Buffer
			pkgmg := clientgen.NewPackages(root, path, g.clientSetAPIPath, string(version.Version), gv.PackageName, &outCommonContent)

			if err := g.writeHeader(&outCommonContent); err != nil {
				root.AddError(err)
			}

			if eachTypeErr := markers.EachType(ctx.Collector, root, func(info *markers.TypeInfo) {
				var outContent bytes.Buffer

				// if not enabled for this type, skip
				if !IsEnabledForMethod(info) {
					return
				}

				/*
					This is the current list of unique genclient markers as of 1.23

					// +genclient
					// +genclient:method=ApplyScale,verb=apply,subresource=scale,input=Scale,result=Scale
					// +genclient:method=ApplyScale,verb=apply,subresource=scale,input=k8s.io/api/autoscaling/v1.Scale,result=k8s.io/api/autoscaling/v1.Scale
					// +genclient:method=CreateScale,verb=create,subresource=scale,input=k8s.io/api/autoscaling/v1.Scale,result=k8s.io/api/autoscaling/v1.Scale
					// +genclient:method=CreateToken,verb=create,subresource=token,input=k8s.io/api/authentication/v1.TokenRequest,result=k8s.io/api/authentication/v1.TokenRequest
					// +genclient:method=GetScale,verb=get,subresource=scale,result=Scale
					// +genclient:method=GetScale,verb=get,subresource=scale,result=k8s.io/api/autoscaling/v1.Scale
					// +genclient:method=UpdateApproval,verb=update,subresource=approval,input=k8s.io/api/certificates/v1.CertificateSigningRequest,result=k8s.io/api/certificates/v1.CertificateSigningRequest
					// +genclient:method=UpdateEphemeralContainers,verb=update,subresource=ephemeralcontainers
					// +genclient:method=UpdateScale,verb=update,subresource=scale,input=Scale,result=Scale
					// +genclient:method=UpdateScale,verb=update,subresource=scale,input=k8s.io/api/autoscaling/v1.Scale,result=k8s.io/api/autoscaling/v1.Scale
					// +genclient:noStatus
					// +genclient:nonNamespaced
					// +genclient:noVerbs
					// +genclient:onlyVerbs=create
					// +genclient:onlyVerbs=create,delete
					// +genclient:readonly
					// +genclient:skipVerbs=deleteCollection
					// +genclient:skipVerbs=get,update
				*/

				var additionalMethods []clientgen.AdditionalMethod
				var skipVerbs []string
				var onlyVerbs []string

				genclientMarkers := info.Markers[GenclientMarker.Name]

				namespaceScoped := info.Markers.Get(nonNamespacedMarker.Name) == nil
				noVerbs := info.Markers.Get(noVerbsMarker.Name) != nil
				hasStatus := HasStatusSubresource(info)
				readOnly := info.Markers.Get(readOnlyMarker.Name) != nil

				// Extract values from skip verbs marker.
				sVerbs := info.Markers.Get(skipVerbsMarker.Name)
				if sVerbs != nil {
					val, ok := sVerbs.(markers.RawArguments)
					if !ok {
						root.AddError(fmt.Errorf("marker defined in wrong format %q", skipVerbsMarker.Name))
					}
					skipVerbs = strings.Split(string(val), ",")
				}

				// Extract values from only verbs marker.
				oVerbs := info.Markers.Get(onlyVerbsMarker.Name)
				if oVerbs != nil {
					val, ok := oVerbs.(markers.RawArguments)
					if !ok {
						root.AddError(fmt.Errorf("marker defined in wrong format %q", onlyVerbsMarker.Name))
					}
					onlyVerbs = append(onlyVerbs, strings.Split(string(val), ",")...)
				}

				if readOnly {
					onlyVerbs = append(onlyVerbs, genutil.ReadonlyVerbs...)
				}

				for _, m := range genclientMarkers {
					gm, ok := m.(genclient)
					if !ok {
						root.AddError(fmt.Errorf("marker defined in wrong format %q", GenclientMarker.Name))
					}

					if gm.Method != nil {
						additionalMethod := clientgen.AdditionalMethod{
							Method:      gm.Method,
							Verb:        gm.Verb,
							Subresource: gm.Subresource,
							Input:       gm.Input,
							Result:      gm.Result,
						}

						if err := verifyAdditionalMethod(additionalMethod); err != nil {
							root.AddError(err)
						}
						additionalMethods = append(additionalMethods, additionalMethod)
					}

					// parse onlyVerbs marker
					if len(onlyVerbs) > 0 {
						onlyVerbsSet := make(map[string]bool)
						for _, o := range onlyVerbs {
							onlyVerbsSet[o] = true
						}

						// Combine the verbs in only Verbs and skip Verbs list.
						skippedverbsList := []string{}
						for _, m := range genutil.SupportedVerbs {
							skip := true
							if _, ok := onlyVerbsSet[m]; ok {
								skip = false
							}

							// Check for conflits if a verb is present in onlyVerb and also in a SkipVerb marker.
							for _, v := range skippedverbsList {
								if v == m {
									root.AddError(fmt.Errorf("verb %q used both in genclient:skipVerbs and genclient:onlyVerbs", v))
								}
							}

							if skip {
								skippedverbsList = append(skippedverbsList, m)
							}
						}
						skipVerbs = skippedverbsList
					}
				}

				a, err := clientgen.NewAPI(
					root,
					info,
					gv.PackageName,
					string(version.Version),
					g.applyConfigGenPkg,
					namespaceScoped,
					additionalMethods,
					&importList,
					skipVerbs,
					onlyVerbs,
					noVerbs,
					hasStatus,
					&outContent,
				)
				if err != nil {
					root.AddError(err)
					return
				}

				err = a.WriteContent()
				if err != nil {
					root.AddError(err)
					return
				}

				outBytes := outContent.Bytes()
				if len(outBytes) > 0 {
					byType[info.Name] = outBytes
				}
			}); eachTypeErr != nil {
				return eachTypeErr
			}

			if len(byType) == 0 {
				continue
			}

			err = pkgmg.WriteContent(&importList)
			if err != nil {
				root.AddError(err)
			}

			var outContent bytes.Buffer
			outContent.Write(outCommonContent.Bytes())
			err = util.WriteMethods(&outContent, byType)
			if err != nil {
				return err
			}

			outBytes := outContent.Bytes()
			formattedBytes, err := format.Source(outBytes)
			if err != nil {
				root.AddError(err)
			} else {
				outBytes = formattedBytes
			}

			filename := gv.Group.PackageName() + string(version.Version) + util.ExtensionGo
			err = util.WriteContent(outBytes, filename, filepath.Join(g.outputDir, g.clientsetName, typedPackageName, gv.Group.PackageName(), string(version.Version)))
			if err != nil {
				root.AddError(err)
				return err
			}
		}
	}
	return util.WriteContent(g.clusterclientWrappers, clientSetFilename, filepath.Join(g.outputDir, g.clientsetName))
}

func verifyAdditionalMethod(m clientgen.AdditionalMethod) error {
	if m.Verb == nil {
		return fmt.Errorf("verb type must be specified (use '// +genclient:method=%s,verb=create')", *m.Method)
	}
	if m.Result != nil {
		supported := false
		for _, v := range util.ResultTypeSupportedVerbs {
			if *m.Verb == v {
				supported = true
				break
			}
		}
		if !supported {
			return fmt.Errorf("%s: result type is not supported for %q verbs (supported verbs: %#v)", *m.Method, *m.Verb, util.ResultTypeSupportedVerbs)
		}
	}

	if m.Input != nil {
		supported := false
		for _, v := range util.InputTypeSupportedVerbs {
			if *m.Verb == v {
				supported = true
				break
			}
		}
		if !supported {
			return fmt.Errorf("%s: input type is not supported for %q verbs (supported verbs: %#v)", *m.Method, *m.Verb, util.InputTypeSupportedVerbs)
		}
	}
	for _, t := range util.UnsupportedExtensionVerbs {
		if *m.Verb == t {
			return fmt.Errorf("verb %q is not supported by extension generator", *m.Verb)
		}
	}
	return nil
}
