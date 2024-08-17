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

package v1

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"

	v1 "acme.corp/pkg/apis/example/v1"
	scheme "acme.corp/pkg/generated/clientset/versioned/scheme"
)

// TestTypesGetter has a method to return a TestTypeInterface.
// A group's client should implement this interface.
type TestTypesGetter interface {
	TestTypes(namespace string) TestTypeInterface
}

// TestTypeInterface has methods to work with TestType resources.
type TestTypeInterface interface {
	Create(ctx context.Context, testType *v1.TestType, opts metav1.CreateOptions) (*v1.TestType, error)
	Update(ctx context.Context, testType *v1.TestType, opts metav1.UpdateOptions) (*v1.TestType, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TestType, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TestTypeList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TestType, err error)
	CreateField(ctx context.Context, testTypeName string, field *v1.Field, opts metav1.CreateOptions) (*v1.Field, error)
	UpdateField(ctx context.Context, testTypeName string, field *v1.Field, opts metav1.UpdateOptions) (*v1.Field, error)
	GetField(ctx context.Context, testTypeName string, options metav1.GetOptions) (*v1.Field, error)

	TestTypeExpansion
}

// testTypes implements TestTypeInterface
type testTypes struct {
	*gentype.ClientWithList[*v1.TestType, *v1.TestTypeList]
}

// newTestTypes returns a TestTypes
func newTestTypes(c *ExampleV1Client, namespace string) *testTypes {
	return &testTypes{
		gentype.NewClientWithList[*v1.TestType, *v1.TestTypeList](
			"testtypes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.TestType { return &v1.TestType{} },
			func() *v1.TestTypeList { return &v1.TestTypeList{} }),
	}
}

// CreateField takes the representation of a field and creates it.  Returns the server's representation of the field, and an error, if there is any.
func (c *testTypes) CreateField(ctx context.Context, testTypeName string, field *v1.Field, opts metav1.CreateOptions) (result *v1.Field, err error) {
	result = &v1.Field{}
	err = c.GetClient().Post().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("field").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(field).
		Do(ctx).
		Into(result)
	return
}

// UpdateField takes the top resource name and the representation of a field and updates it. Returns the server's representation of the field, and an error, if there is any.
func (c *testTypes) UpdateField(ctx context.Context, testTypeName string, field *v1.Field, opts metav1.UpdateOptions) (result *v1.Field, err error) {
	result = &v1.Field{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("field").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(field).
		Do(ctx).
		Into(result)
	return
}

// GetField takes name of the testType, and returns the corresponding v1.Field object, and an error if there is any.
func (c *testTypes) GetField(ctx context.Context, testTypeName string, options metav1.GetOptions) (result *v1.Field, err error) {
	result = &v1.Field{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("testtypes").
		Name(testTypeName).
		SubResource("field").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}