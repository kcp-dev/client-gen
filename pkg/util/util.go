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

package util

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"k8s.io/klog/v2"
	"sigs.k8s.io/controller-tools/pkg/genall"
)

// LowerFirst sets the first alphabet to lowerCase.
func LowerFirst(s string) string {
	return strings.ToLower(string(s[0])) + s[1:]
}

// UpperFirst sets the first alphabet to upperCase/
func UpperFirst(s string) string {
	return strings.ToUpper(string(s[0])) + s[1:]
}

type Generator interface {
	WriteContent(w io.Writer) error
}

func WriteGeneratedCode(ctx *genall.GenerationContext, header string, generator Generator, intoPath string) error {
	data := &bytes.Buffer{}
	if n, err := data.WriteString(header); err != nil {
		return err
	} else if n != len(header) {
		return io.ErrShortWrite
	}

	if err := generator.WriteContent(data); err != nil {
		return err
	}

	if err := os.MkdirAll(filepath.Dir(intoPath), os.ModePerm); err != nil {
		return fmt.Errorf("couldn't create directory: %w", err)
	}

	outputFile, err := ctx.Open(nil, intoPath)
	if err != nil {
		return fmt.Errorf("failed to open: %w", err)
	}
	defer func() {
		if err := outputFile.Close(); err != nil {
			klog.Background().Error(err, "failed to close output file")
		}
	}()
	dataBytes := data.Bytes()
	n, err := outputFile.Write(dataBytes)
	if err != nil {
		return err
	}
	if n < len(dataBytes) {
		return io.ErrShortWrite
	}

	return nil
}
