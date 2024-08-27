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
	v1beta2 "k8s.io/api/apps/v1beta2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	appsv1beta2 "k8s.io/client-go/applyconfigurations/apps/v1beta2"
	upstreamappsv1beta2client "k8s.io/client-go/kubernetes/typed/apps/v1beta2"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/apps/v1beta2"
)

var daemonsetsResource = v1beta2.SchemeGroupVersion.WithResource("daemonsets")

var daemonsetsKind = v1beta2.SchemeGroupVersion.WithKind("DaemonSet")

// daemonSetsClusterClient implements daemonSetInterface
type daemonSetsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *daemonSetsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.DaemonSetNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &daemonSetsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of DaemonSets that match those selectors.
func (c *daemonSetsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.DaemonSetList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(daemonsetsResource, daemonsetsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta2.DaemonSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.DaemonSetList{ListMeta: obj.(*v1beta2.DaemonSetList).ListMeta}
	for _, item := range obj.(*v1beta2.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested daemonSets across all clusters.
func (c *daemonSetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(daemonsetsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type daemonSetsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *daemonSetsNamespacer) Namespace(namespace string) upstreamappsv1beta2client.DaemonSetInterface {
	return &daemonSetsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type daemonSetsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *daemonSetsClient) Create(ctx context.Context, daemonSet *v1beta2.DaemonSet, opts metav1.CreateOptions) (*v1beta2.DaemonSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(daemonsetsResource, c.ClusterPath, c.Namespace, daemonSet), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}

func (c *daemonSetsClient) Update(ctx context.Context, daemonSet *v1beta2.DaemonSet, opts metav1.UpdateOptions) (*v1beta2.DaemonSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(daemonsetsResource, c.ClusterPath, c.Namespace, daemonSet), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}

func (c *daemonSetsClient) UpdateStatus(ctx context.Context, daemonSet *v1beta2.DaemonSet, opts metav1.UpdateOptions) (*v1beta2.DaemonSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(daemonsetsResource, c.ClusterPath, "status", c.Namespace, daemonSet), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}

func (c *daemonSetsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(daemonsetsResource, c.ClusterPath, c.Namespace, name, opts), &v1beta2.DaemonSet{})
	return err
}

func (c *daemonSetsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(daemonsetsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.DaemonSetList{})
	return err
}

func (c *daemonSetsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta2.DaemonSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(daemonsetsResource, c.ClusterPath, c.Namespace, name), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}

// List takes label and field selectors, and returns the list of v1beta2.DaemonSet that match those selectors.
func (c *daemonSetsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta2.DaemonSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(daemonsetsResource, daemonsetsKind, c.ClusterPath, c.Namespace, opts), &v1beta2.DaemonSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.DaemonSetList{ListMeta: obj.(*v1beta2.DaemonSetList).ListMeta}
	for _, item := range obj.(*v1beta2.DaemonSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *daemonSetsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(daemonsetsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *daemonSetsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta2.DaemonSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(daemonsetsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}

func (c *daemonSetsClient) Apply(ctx context.Context, applyConfiguration *appsv1beta2.DaemonSetApplyConfiguration, opts metav1.ApplyOptions) (*v1beta2.DaemonSet, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(daemonsetsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta2.DaemonSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.DaemonSet), err
}
