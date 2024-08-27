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

var replicasetsResource = v1beta2.SchemeGroupVersion.WithResource("replicasets")

var replicasetsKind = v1beta2.SchemeGroupVersion.WithKind("ReplicaSet")

// replicaSetsClusterClient implements replicaSetInterface
type replicaSetsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *replicaSetsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.ReplicaSetNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &replicaSetsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of ReplicaSets that match those selectors.
func (c *replicaSetsClusterClient) List(ctx context.Context, opts v1.ListOptions) (result *v1beta2.ReplicaSetList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicasetsResource, replicasetsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1beta2.ReplicaSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ReplicaSetList{ListMeta: obj.(*v1beta2.ReplicaSetList).ListMeta}
	for _, item := range obj.(*v1beta2.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested replicaSets across all clusters.
func (c *replicaSetsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicasetsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type replicaSetsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *replicaSetsNamespacer) Namespace(namespace string) upstreamappsv1beta2client.ReplicaSetInterface {
	return &replicaSetsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type replicaSetsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *replicaSetsClient) Create(ctx context.Context, replicaSet *v1beta2.ReplicaSet, opts metav1.CreateOptions) (*v1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(replicasetsResource, c.ClusterPath, c.Namespace, replicaSet), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Update(ctx context.Context, replicaSet *v1beta2.ReplicaSet, opts metav1.UpdateOptions) (*v1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(replicasetsResource, c.ClusterPath, c.Namespace, replicaSet), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) UpdateStatus(ctx context.Context, replicaSet *v1beta2.ReplicaSet, opts metav1.UpdateOptions) (*v1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(replicasetsResource, c.ClusterPath, "status", c.Namespace, replicaSet), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(replicasetsResource, c.ClusterPath, c.Namespace, name, opts), &v1beta2.ReplicaSet{})
	return err
}

func (c *replicaSetsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(replicasetsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta2.ReplicaSetList{})
	return err
}

func (c *replicaSetsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(replicasetsResource, c.ClusterPath, c.Namespace, name), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}

// List takes label and field selectors, and returns the list of v1beta2.ReplicaSet that match those selectors.
func (c *replicaSetsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1beta2.ReplicaSetList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(replicasetsResource, replicasetsKind, c.ClusterPath, c.Namespace, opts), &v1beta2.ReplicaSetList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta2.ReplicaSetList{ListMeta: obj.(*v1beta2.ReplicaSetList).ListMeta}
	for _, item := range obj.(*v1beta2.ReplicaSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *replicaSetsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(replicasetsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *replicaSetsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1beta2.ReplicaSet, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicasetsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}

func (c *replicaSetsClient) Apply(ctx context.Context, applyConfiguration *appsv1beta2.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (*v1beta2.ReplicaSet, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(replicasetsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1beta2.ReplicaSet{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta2.ReplicaSet), err
}
