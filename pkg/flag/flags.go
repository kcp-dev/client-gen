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

package flag

import (
	"errors"

	"github.com/spf13/pflag"
)

// Flags - Options accepted by generator
type Flags struct {
	// OutputDir is where the generated code is to be written to.
	OutputDir string
	// InputDir is path to the input APIs (types.go)
	InputDir string
	// ClientsetAPIPath is the path to where client sets are scaffolded by codegen.
	ClientsetAPIPath string
	// List of group versions for which the wrappers are to be generated.
	GroupVersions []string
	// Path to the headerfile.
	GoHeaderFilePath string
	// ClientsetName is the name of the clientset to be generated.
	ClientsetName string
}

func (f *Flags) AddTo(flagset *pflag.FlagSet) {
	// TODO: Figure out if its worth defaulting it to pkg/api/...
	flagset.StringVar(&f.InputDir, "input-dir", "", "Input directory where types are defined. It is assumed that 'types.go' is present inside <InputDir>/pkg/apis.")
	flagset.StringVar(&f.OutputDir, "output-dir", "output", "Output directory where wrapped clients will be generated. The wrappers will be present in '<output-dir>/generated' path.")
	flagset.StringVar(&f.ClientsetAPIPath, "clientset-api-path", "/apis", "package path where clients are generated.")

	flagset.StringArrayVar(&f.GroupVersions, "group-versions", []string{}, "specify group versions for the clients.")
	flagset.StringVar(&f.GoHeaderFilePath, "go-header-file", "", "path to headerfile for the generated text.")
	flagset.StringVar(&f.ClientsetName, "clientset-name", "clientset", "the name of the generated clientset package.")
}

// ValidateFlags checks if the inputs provided through flags are valid and
// if so, sets defaults.
// TODO: Remove this and bind options to all the generators, see
// https://github.com/kcp-dev/code-generator/issues/4
func ValidateFlags(f Flags) error {
	if f.InputDir == "" {
		return errors.New("input path to API definition is required.")
	}

	if f.ClientsetAPIPath == "" {
		return errors.New("specifying client API path is required currently.")
	}

	if len(f.GroupVersions) == 0 {
		return errors.New("list of group versions for which the clients are to be generated is required.")
	}

	return nil
}
