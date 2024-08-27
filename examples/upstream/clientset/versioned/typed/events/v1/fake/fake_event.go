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
	v1 "k8s.io/api/events/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	eventsv1 "k8s.io/client-go/applyconfigurations/events/v1"
	upstreameventsv1client "k8s.io/client-go/kubernetes/typed/events/v1"
	"k8s.io/client-go/testing"
	kcp "k8s.io/code-generator/examples/upstream/clientset/versioned/typed/events/v1"
)

var eventsResource = v1.SchemeGroupVersion.WithResource("events")

var eventsKind = v1.SchemeGroupVersion.WithKind("Event")

// eventsClusterClient implements eventInterface
type eventsClusterClient struct {
	*kcptesting.Fake
}

// Cluster scopes the client down to a particular cluster.
func (c *eventsClusterClient) Cluster(clusterPath logicalcluster.Path) kcp.EventNamespacer {
	if clusterPath == logicalcluster.Wildcard {
		panic("A specific cluster must be provided when scoping, not the wildcard.")
	}

	return &eventsNamespacer{Fake: c.Fake, ClusterPath: clusterPath}
}

// List takes label and field selectors, and returns the list of Events that match those selectors.
func (c *eventsClusterClient) List(ctx context.Context, opts metav1.ListOptions) (result *v1.EventList, err error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(eventsResource, eventsKind, logicalcluster.Wildcard, metav1.NamespaceAll, opts), &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EventList{ListMeta: obj.(*v1.EventList).ListMeta}
	for _, item := range obj.(*v1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested events across all clusters.
func (c *eventsClusterClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(eventsResource, logicalcluster.Wildcard, metav1.NamespaceAll, opts))
}

type eventsNamespacer struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
}

func (n *eventsNamespacer) Namespace(namespace string) upstreameventsv1client.EventInterface {
	return &eventsClient{Fake: n.Fake, ClusterPath: n.ClusterPath, Namespace: namespace}
}

type eventsClient struct {
	*kcptesting.Fake
	ClusterPath logicalcluster.Path
	Namespace   string
}

func (c *eventsClient) Create(ctx context.Context, event *v1.Event, opts metav1.CreateOptions) (*v1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace, event), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

func (c *eventsClient) Update(ctx context.Context, event *v1.Event, opts metav1.UpdateOptions) (*v1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace, event), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

func (c *eventsClient) UpdateStatus(ctx context.Context, event *v1.Event, opts metav1.UpdateOptions) (*v1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewUpdateSubresourceAction(eventsResource, c.ClusterPath, "status", c.Namespace, event), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

func (c *eventsClient) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	_, err := c.Fake.Invokes(kcptesting.NewDeleteActionWithOptions(eventsResource, c.ClusterPath, c.Namespace, name, opts), &v1.Event{})
	return err
}

func (c *eventsClient) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	action := kcptesting.NewDeleteCollectionAction(eventsResource, c.ClusterPath, c.Namespace, listOpts)

	_, err := c.Fake.Invokes(action, &v1.EventList{})
	return err
}

func (c *eventsClient) Get(ctx context.Context, name string, options metav1.GetOptions) (*v1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewGetAction(eventsResource, c.ClusterPath, c.Namespace, name), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

// List takes label and field selectors, and returns the list of v1.Event that match those selectors.
func (c *eventsClient) List(ctx context.Context, opts metav1.ListOptions) (*v1.EventList, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace, opts), &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1.EventList{ListMeta: obj.(*v1.EventList).ListMeta}
	for _, item := range obj.(*v1.EventList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

func (c *eventsClient) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	return c.Fake.InvokesWatch(kcptesting.NewWatchAction(eventsResource, c.ClusterPath, c.Namespace, opts))
}

func (c *eventsClient) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (*v1.Event, error) {
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(eventsResource, c.ClusterPath, c.Namespace, name, pt, data, subresources...), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

func (c *eventsClient) Apply(ctx context.Context, applyConfiguration *eventsv1.EventApplyConfiguration, opts metav1.ApplyOptions) (*v1.Event, error) {
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
	obj, err := c.Fake.Invokes(kcptesting.NewPatchSubresourceAction(eventsResource, c.ClusterPath, c.Namespace, *name, types.ApplyPatchType, data), &v1.Event{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1.Event), err
}

func (c *eventsClient) CreateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootCreateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = core.NewCreateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Update replaces an existing event. Returns the copy of the event the server returns, or an error.
func (c *eventsClient) UpdateWithEventNamespace(event *v1.Event) (*v1.Event, error) {
	action := core.NewRootUpdateAction(eventsResource, c.ClusterPath, event)
	if c.Namespace != "" {
		action = core.NewUpdateAction(eventsResource, c.ClusterPath, c.Namespace, event)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// PatchWithEventNamespace patches an existing event. Returns the copy of the event the server returns, or an error.
// TODO: Should take a PatchType as an argument probably.
func (c *eventsClient) PatchWithEventNamespace(event *v1.Event, data []byte) (*v1.Event, error) {
	// TODO: Should be configurable to support additional patch strategies.
	pt := types.StrategicMergePatchType
	action := core.NewRootPatchAction(eventsResource, c.ClusterPath, event.Name, pt, data)
	if c.Namespace != "" {
		action = core.NewPatchAction(eventsResource, c.ClusterPath, c.Namespace, event.Name, pt, data)
	}
	obj, err := c.Fake.Invokes(action, event)
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.Event), err
}

// Search returns a list of events matching the specified object.
func (c *eventsClient) Search(scheme *runtime.Scheme, objOrRef runtime.Object) (*v1.EventList, error) {
	action := core.NewRootListAction(eventsResource, eventsKind, c.ClusterPath, metav1.ListOptions{})
	if c.Namespace != "" {
		action = core.NewListAction(eventsResource, eventsKind, c.ClusterPath, c.Namespace, metav1.ListOptions{})
	}
	obj, err := c.Fake.Invokes(action, &v1.EventList{})
	if obj == nil {
		return nil, err
	}

	return obj.(*v1.EventList), err
}

func (c *eventsClient) GetFieldSelector(involvedObjectName, involvedObjectNamespace, involvedObjectKind, involvedObjectUID *string) fields.Selector {
	action := core.GenericActionImpl{}
	action.Verb = "get-field-selector"
	action.Resource = eventsResource
	action.ClusterPath = c.ClusterPath

	_, _ = c.Fake.Invokes(action, nil)
	return fields.Everything()
}
