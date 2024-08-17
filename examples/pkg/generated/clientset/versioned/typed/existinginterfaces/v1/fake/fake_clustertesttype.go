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

// Code generated by client-gen-v0.31. DO NOT EDIT.

package fake

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1 "acme.corp/pkg/apis/existinginterfaces/v1"
)

// FakeClusterTestTypes implements ClusterTestTypeInterface
type FakeClusterTestTypes struct {
	Fake *FakeExistinginterfacesV1
}

var clustertesttypesResource = v1.SchemeGroupVersion.WithResource("clustertesttypes")

var clustertesttypesKind = v1.SchemeGroupVersion.WithKind("ClusterTestType")

// Get takes name of the clusterTestType, and returns the corresponding clusterTestType object, and an error if there is any.
func (c *FakeClusterTestTypes) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.ClusterTestType, err error) {
	emptyResult := &v1.ClusterTestType{}
	obj, err := c.Fake.
		Invokes(testing.NewRootGetActionWithOptions(clustertesttypesResource, name, options), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterTestType), err
}

// List takes label and field selectors, and returns the list of ClusterTestTypes that match those selectors.
func (c *FakeClusterTestTypes) List(ctx context.Context, opts metav1.ListOptions) (result *v1.ClusterTestTypeList, err error) {
	emptyResult := &v1.ClusterTestTypeList{}
	obj, err := c.Fake.
		Invokes(testing.NewRootListActionWithOptions(clustertesttypesResource, clustertesttypesKind, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.ClusterTestTypeList{ListMeta: obj.(*v1.ClusterTestTypeList).ListMeta}
	for _, item := range obj.(*v1.ClusterTestTypeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterTestTypes.
func (c *FakeClusterTestTypes) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchActionWithOptions(clustertesttypesResource, opts))
}

// Create takes the representation of a clusterTestType and creates it.  Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *FakeClusterTestTypes) Create(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.CreateOptions) (result *v1.ClusterTestType, err error) {
	emptyResult := &v1.ClusterTestType{}
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateActionWithOptions(clustertesttypesResource, clusterTestType, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterTestType), err
}

// Update takes the representation of a clusterTestType and updates it. Returns the server's representation of the clusterTestType, and an error, if there is any.
func (c *FakeClusterTestTypes) Update(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (result *v1.ClusterTestType, err error) {
	emptyResult := &v1.ClusterTestType{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateActionWithOptions(clustertesttypesResource, clusterTestType, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterTestType), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterTestTypes) UpdateStatus(ctx context.Context, clusterTestType *v1.ClusterTestType, opts metav1.UpdateOptions) (result *v1.ClusterTestType, err error) {
	emptyResult := &v1.ClusterTestType{}
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceActionWithOptions(clustertesttypesResource, "status", clusterTestType, opts), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterTestType), err
}

// Delete takes name of the clusterTestType and deletes it. Returns an error if one occurs.
func (c *FakeClusterTestTypes) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clustertesttypesResource, name, opts), &v1.ClusterTestType{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterTestTypes) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := testing.NewRootDeleteCollectionActionWithOptions(clustertesttypesResource, opts, listOpts)

	_, err := c.Fake.Invokes(action, &v1.ClusterTestTypeList{})
	return err
}

// Patch applies the patch and returns the patched clusterTestType.
func (c *FakeClusterTestTypes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ClusterTestType, err error) {
	emptyResult := &v1.ClusterTestType{}
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceActionWithOptions(clustertesttypesResource, name, pt, data, opts, subresources...), emptyResult)
	if obj == nil {
		return emptyResult, err
	}
	return obj.(*v1.ClusterTestType), err
}
