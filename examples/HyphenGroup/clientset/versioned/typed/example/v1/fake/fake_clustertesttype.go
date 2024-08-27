/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	v1 "k8s.io/code-generator/examples/HyphenGroup/apis/example/v1"
)

var clustertesttypesResource = v1.SchemeGroupVersion.WithResource("clustertesttypes")

var clustertesttypesKind = v1.SchemeGroupVersion.WithKind("ClusterTestType")

// clusterTestTypesClusterClient implements clusterTestTypeInterface
type clusterTestTypesClusterClient struct {
	*kcptesting.Fake
}
