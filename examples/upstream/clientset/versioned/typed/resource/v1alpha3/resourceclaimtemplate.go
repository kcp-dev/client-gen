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

package v1alpha3

import (
	"context"

	v1alpha3 "k8s.io/api/resource/v1alpha3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	resourcev1alpha3 "k8s.io/code-generator/examples/upstream/applyconfiguration/resource/v1alpha3"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ResourceClaimTemplatesGetter has a method to return a ResourceClaimTemplateInterface.
// A group's client should implement this interface.
type ResourceClaimTemplatesGetter interface {
	ResourceClaimTemplates(namespace string) ResourceClaimTemplateInterface
}

// ResourceClaimTemplateInterface has methods to work with ResourceClaimTemplate resources.
type ResourceClaimTemplateInterface interface {
	Create(ctx context.Context, resourceClaimTemplate *v1alpha3.ResourceClaimTemplate, opts v1.CreateOptions) (*v1alpha3.ResourceClaimTemplate, error)
	Update(ctx context.Context, resourceClaimTemplate *v1alpha3.ResourceClaimTemplate, opts v1.UpdateOptions) (*v1alpha3.ResourceClaimTemplate, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha3.ResourceClaimTemplate, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha3.ResourceClaimTemplateList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.ResourceClaimTemplate, err error)
	Apply(ctx context.Context, resourceClaimTemplate *resourcev1alpha3.ResourceClaimTemplateApplyConfiguration, opts v1.ApplyOptions) (result *v1alpha3.ResourceClaimTemplate, err error)
	ResourceClaimTemplateExpansion
}

// resourceClaimTemplates implements ResourceClaimTemplateInterface
type resourceClaimTemplates struct {
	*gentype.ClientWithListAndApply[*v1alpha3.ResourceClaimTemplate, *v1alpha3.ResourceClaimTemplateList, *resourcev1alpha3.ResourceClaimTemplateApplyConfiguration]
}

// newResourceClaimTemplates returns a ResourceClaimTemplates
func newResourceClaimTemplates(c *ResourceV1alpha3Client, namespace string) *resourceClaimTemplates {
	return &resourceClaimTemplates{
		gentype.NewClientWithListAndApply[*v1alpha3.ResourceClaimTemplate, *v1alpha3.ResourceClaimTemplateList, *resourcev1alpha3.ResourceClaimTemplateApplyConfiguration](
			"resourceclaimtemplates",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1alpha3.ResourceClaimTemplate { return &v1alpha3.ResourceClaimTemplate{} },
			func() *v1alpha3.ResourceClaimTemplateList { return &v1alpha3.ResourceClaimTemplateList{} }),
	}
}