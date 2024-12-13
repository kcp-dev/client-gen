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

package listergen

import (
	"path/filepath"
	"strings"

	"k8s.io/gengo/v2/namer"
	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
	"sigs.k8s.io/controller-tools/pkg/markers"

	"github.com/kcp-dev/code-generator/v2/pkg/internal/listergen"
	"github.com/kcp-dev/code-generator/v2/pkg/parser"
	"github.com/kcp-dev/code-generator/v2/pkg/util"
)

type Generator struct {
	// Name is the name of this client-set, e.g. "kubernetes"
	Name string `marker:",optional"`

	// HeaderFile specifies the header text (e.g. license) to prepend to generated files.
	HeaderFile string `marker:",optional"`

	// Year specifies the year to substitute for " YEAR" in the header file.
	Year string `marker:",optional"`

	// APIPackagePath is the root directory under which API types exist.
	// e.g. "k8s.io/api"
	APIPackagePath string `marker:"apiPackagePath"`

	// SingleClusterListerPackagePath is the root directory under which single-cluster-aware listers exist,
	// for the case where we're only generating new code "on top" to enable multi-cluster use-cases.
	// e.g. "k8s.io/client-go/listers"
	SingleClusterListerPackagePath string `marker:",optional"`
}

func (Generator) RegisterMarkers(into *markers.Registry) error {
	return markers.RegisterAll(into,
		parser.GenclientMarker(),
		parser.NonNamespacedMarker(),
		parser.GroupNameMarker(),
		parser.NoVerbsMarker(),
		parser.ReadOnlyMarker(),
		parser.SkipVerbsMarker(),
		parser.OnlyVerbsMarker(),
	)
}

// Generate will generate listers for all types that have generated clients and support LIST + WATCH verbs.
func (g Generator) Generate(ctx *genall.GenerationContext) error {
	var headerText string

	if g.HeaderFile != "" {
		headerBytes, err := ctx.ReadFile(g.HeaderFile)
		if err != nil {
			return err
		}
		headerText = string(headerBytes)
	}
	var replacement string
	if g.Year != "" {
		replacement = " " + g.Year
	}
	headerText = strings.ReplaceAll(headerText, " YEAR", replacement)

	groupVersionKinds, err := parser.CollectKinds(ctx, "list", "watch")
	if err != nil {
		return err
	}

	for group, versions := range groupVersionKinds {
		for version, kinds := range versions {
			groupInfo := toGroupVersionInfo(group, version)
			for _, kind := range kinds {
				listerDir := filepath.Join("listers", group.PackageName(), version.PackageName())
				outputFile := filepath.Join(listerDir, strings.ToLower(kind.String())+".go")
				logger := klog.Background().WithValues(
					"group", group.String(),
					"version", version.String(),
					"kind", kind.String(),
					"path", outputFile,
				)
				logger.Info("generating lister")

				if err := util.WriteGeneratedCode(ctx, headerText, &listergen.Lister{
					Group:                          groupInfo,
					APIPackagePath:                 g.APIPackagePath,
					Kind:                           kind,
					SingleClusterListerPackagePath: g.SingleClusterListerPackagePath,
				}, outputFile); err != nil {
					return err
				}

				outputFile = filepath.Join(listerDir, strings.ToLower(kind.String())+"_expansion.go")
				logger = logger.WithValues(
					"path", outputFile,
				)
				logger.Info("generating lister expansion")

				if err := util.InitializeGeneratedCode(ctx, headerText, &listergen.Expansions{
					Group:                 groupInfo,
					Kind:                  kind,
					UseUpstreamInterfaces: g.SingleClusterListerPackagePath != "",
				}, outputFile); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// adapted from https://github.com/kubernetes/kubernetes/blob/8f269d6df2a57544b73d5ca35e04451373ef334c/staging/src/k8s.io/code-generator/cmd/client-gen/types/helpers.go#L87-L103
func toGroupVersionInfo(group parser.Group, version parser.PackageVersion) parser.Group {
	return parser.Group{
		Group:                group.Group,
		Version:              parser.Version(namer.IC(version.Version.String())),
		PackageAlias:         strings.ReplaceAll(strings.ToLower(group.GoName+version.Version.NonEmpty()), "-", ""),
		GoName:               group.GoName,
		LowerCaseGroupGoName: namer.IL(group.GoName),
	}
}
