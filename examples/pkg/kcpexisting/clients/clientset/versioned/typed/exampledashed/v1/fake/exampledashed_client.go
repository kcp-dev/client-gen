//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package fake

import (
	"github.com/kcp-dev/logicalcluster/v3"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"k8s.io/client-go/rest"

	exampledashedv1 "acme.corp/pkg/generated/clientset/versioned/typed/exampledashed/v1"
	kcpexampledashedv1 "acme.corp/pkg/kcpexisting/clients/clientset/versioned/typed/exampledashed/v1"
)

var _ kcpexampledashedv1.ExampleDashedV1ClusterInterface = (*ExampleDashedV1ClusterClient)(nil)

type ExampleDashedV1ClusterClient struct {
	*kcptesting.Fake
}

func (c *ExampleDashedV1ClusterClient) Cluster(clusterPath logicalcluster.Path) exampledashedv1.ExampleDashedV1Interface {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}
	return &ExampleDashedV1Client{Fake: c.Fake, ClusterPath: clusterPath}
}

func (c *ExampleDashedV1ClusterClient) TestTypes() kcpexampledashedv1.TestTypeClusterInterface {
	return &testTypesClusterClient{Fake: c.Fake}
}

func (c *ExampleDashedV1ClusterClient) ClusterTestTypes() kcpexampledashedv1.ClusterTestTypeClusterInterface {
	return &clusterTestTypesClusterClient{Fake: c.Fake}
}

var _ exampledashedv1.ExampleDashedV1Interface = (*ExampleDashedV1Client)(nil)

type ExampleDashedV1Client struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (c *ExampleDashedV1Client) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}

func (c *ExampleDashedV1Client) TestTypes(namespace string) exampledashedv1.TestTypeInterface {
	return &testTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath, Namespace: namespace}
}

func (c *ExampleDashedV1Client) ClusterTestTypes() exampledashedv1.ClusterTestTypeInterface {
	return &clusterTestTypesClient{Fake: c.Fake, ClusterPath: c.ClusterPath}
}
