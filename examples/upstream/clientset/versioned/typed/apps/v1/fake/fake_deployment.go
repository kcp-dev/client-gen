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
	"context"
	json "encoding/json"
	"fmt"

	kcptesting "github.com/kcp-dev/client-go/third_party/k8s.io/client-go/testing"
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	appsv1 "k8s.io/client-go/applyconfigurations/apps/v1"
	upstreamappsv1client "k8s.io/client-go/kubernetes/typed/apps/v1"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/apps/v1"
)

var deploymentsResource = v1.SchemeGroupVersion.WithResource("deployments")

var deploymentsKind = v1.SchemeGroupVersion.WithKind("Deployment")

// deploymentsClusterClient implements deploymentInterface
type deploymentsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *deploymentsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.DeploymentNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &deploymentsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Deployments that match those selectors.
func (c *deploymentsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.DeploymentList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DeploymentList{ListMeta: obj.(*v1.DeploymentList).ListMeta}
	for _, item := range obj.(*v1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested deployments across all clusters.
func (c *deploymentsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type deploymentsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *deploymentsNamespacer) Namespace(namespace string) upstreamappsv1client.DeploymentInterface {
	return &deploymentsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type deploymentsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *deploymentsClient) Create(ctx context.Context, deployment *v1.Deployment, opts metav1.CreateOptions) (*v1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(deploymentsResource, c.ClusterPath, c.Namespace, deployment), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}

func (c *deploymentsClient) Update(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (*v1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(deploymentsResource, c.ClusterPath, c.Namespace, deployment), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}

func (c *deploymentsClient) UpdateStatus(ctx context.Context, deployment *v1.Deployment, opts metav1.UpdateOptions) (*v1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(deploymentsResource, c.ClusterPath, "status", c.Namespace, deployment), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}

func (c *deploymentsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(deploymentsResource, c.ClusterPath, c.Namespace, name, opts), &v1.Deployment{})
	return err
}

func (c *deploymentsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(deploymentsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.DeploymentList{})
	return err
}

func (c *deploymentsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(deploymentsResource, c.ClusterPath, c.Namespace, name), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}

// List takes label and field selectors, and returns the list of v1.Deployment that match those selectors.
func (c *deploymentsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.DeploymentList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(deploymentsResource, deploymentsKind, c.ClusterPath, c.Namespace, opts), &v1.DeploymentList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.DeploymentList{ListMeta: obj.(*v1.DeploymentList).ListMeta}
	for _, item := range obj.(*v1.DeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *deploymentsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(deploymentsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *deploymentsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.Deployment, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}

func (c *deploymentsClient) Apply(ctx context.Context, applyConfiguration *appsv1.DeploymentApplyConfiguration, opts metav1.ApplyOptions) (*v1.Deployment, error) {
	if applyConfiguration == nil {
		return nil, fmt.Errorf("applyConfiguration provided to Apply must not be nil")
	}
	data, err := json.Marshal(applyConfiguration)
	if err != nil {
		return nil, err
	}
	name := applyConfiguration.Name
	if name == nil {
		return nil, fmt.Errorf("applyConfiguration.Name must be provided to Apply")
	}
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(deploymentsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1.Deployment{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Deployment), err
}
