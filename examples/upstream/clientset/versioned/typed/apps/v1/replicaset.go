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

package v1

import (
	"context"
	json "encoding/json"
	"fmt"

	v1 "k8s.io/api/apps/v1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
	appsv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/apps/v1"
	applyconfigurationautoscalingv1 "k8s.io/code-generator/examples/upstream/applyconfiguration/autoscaling/v1"
	scheme "k8s.io/code-generator/examples/upstream/clientset/versioned/scheme"
)

// ReplicaSetsGetter has a method to return a ReplicaSetInterface.
// A group's client should implement this interface.
type ReplicaSetsGetter interface {
	ReplicaSets(namespace string) ReplicaSetInterface
}

// ReplicaSetInterface has methods to work with ReplicaSet resources.
type ReplicaSetInterface interface {
	Create(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.CreateOptions) (*v1.ReplicaSet, error)
	Update(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (*v1.ReplicaSet, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, replicaSet *v1.ReplicaSet, opts metav1.UpdateOptions) (*v1.ReplicaSet, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.ReplicaSet, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.ReplicaSetList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.ReplicaSet, err error)
	Apply(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, replicaSet *appsv1.ReplicaSetApplyConfiguration, opts metav1.ApplyOptions) (result *v1.ReplicaSet, err error)
	GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (*autoscalingv1.Scale, error)
	ApplyScale(ctx context.Context, replicaSetName string, scale *applyconfigurationautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (*autoscalingv1.Scale, error)

	ReplicaSetExpansion
}

// replicaSets implements ReplicaSetInterface
type replicaSets struct {
	*gentype.ClientWithListAndApply[*v1.ReplicaSet, *v1.ReplicaSetList, *appsv1.ReplicaSetApplyConfiguration]
}

// newReplicaSets returns a ReplicaSets
func newReplicaSets(c *AppsV1Client, namespace string) *replicaSets {
	return &replicaSets{
		gentype.NewClientWithListAndApply[*v1.ReplicaSet, *v1.ReplicaSetList, *appsv1.ReplicaSetApplyConfiguration](
			"replicasets",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *v1.ReplicaSet { return &v1.ReplicaSet{} },
			func() *v1.ReplicaSetList { return &v1.ReplicaSetList{} }),
	}
}

// GetScale takes name of the replicaSet, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *replicaSets) GetScale(ctx context.Context, replicaSetName string, options metav1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Get().
		Namespace(c.GetNamespace()).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *replicaSets) UpdateScale(ctx context.Context, replicaSetName string, scale *autoscalingv1.Scale, opts metav1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.GetClient().Put().
		Namespace(c.GetNamespace()).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}

// ApplyScale takes top resource name and the apply declarative configuration for scale,
// applies it and returns the applied scale, and an error, if there is any.
func (c *replicaSets) ApplyScale(ctx context.Context, replicaSetName string, scale *applyconfigurationautoscalingv1.ScaleApplyConfiguration, opts metav1.ApplyOptions) (result *autoscalingv1.Scale, err error) {
	if scale == nil {
		return nil, fmt.Errorf("scale provided to ApplyScale must not be nil")
	}
	patchOpts := opts.ToPatchOptions()
	data, err := json.Marshal(scale)
	if err != nil {
		return nil, err
	}

	result = &autoscalingv1.Scale{}
	err = c.GetClient().Patch(types.ApplyPatchType).
		Namespace(c.GetNamespace()).
		Resource("replicasets").
		Name(replicaSetName).
		SubResource("scale").
		VersionedParams(&patchOpts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}