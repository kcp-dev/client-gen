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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1 "k8s.io/api/discovery/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EndpointSliceLister helps list EndpointSlices.
// All objects returned here must be treated as read-only.
type EndpointSliceLister interface {
	// List lists all EndpointSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EndpointSlice, err error)
	// EndpointSlices returns an object that can list and get EndpointSlices.
	EndpointSlices(namespace string) EndpointSliceNamespaceLister
	EndpointSliceListerExpansion
}

// EndpointSliceClusterLister helps list EndpointSlices.
// All objects returned here must be treated as read-only.
type EndpointSliceClusterLister interface {
	// List lists all EndpointSlices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EndpointSlice, err error)
	EndpointSliceClusterListerExpansion
}

// endpointSliceLister implements the EndpointSliceLister interface.
type endpointSliceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// endpointSliceLister implements the EndpointSliceClusterLister interface.
type endpointSliceClusterLister struct {
	indexer cache.Indexer
}

// List lists all EndpointSlices in the indexer.
func (s *endpointSliceLister) List(selector labels.Selector) (ret []*v1.EndpointSlice, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.EndpointSlice))
	})
	return ret, err
}

// Get retrieves the  EndpointSlice from the indexer for a given workspace, namespace and name.
func (s endpointSliceLister) Get(name string) (*v1.EndpointSlice, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("endpointslice"), name)
	}
	return obj.(*v1.EndpointSlice), nil
}

// NewEndpointSliceClusterLister returns a new EndpointSliceClusterLister.
func NewEndpointSliceClusterLister(indexer cache.Indexer) EndpointSliceClusterLister {
	return &endpointSliceClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get EndpointSlice.
func (s *endpointSliceClusterLister) Cluster(clusterName logicalcluster.Name) EndpointSliceLister {
	return &endpointSliceLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all EndpointSlices in the indexer.
func (s *endpointSliceClusterLister) List(selector labels.Selector) (ret []*v1.EndpointSlice, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.EndpointSlice))
	})
	return ret, err
}

// EndpointSlices returns an object that can list and get EndpointSlices.
func (s *endpointSliceLister) EndpointSlices(namespace string) EndpointSliceNamespaceLister {
	return endpointSliceNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// EndpointSliceNamespaceLister helps list and get EndpointSlices.
// All objects returned here must be treated as read-only.
type EndpointSliceNamespaceLister interface {
	// List lists all EndpointSlices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.EndpointSlice, err error)
	// Get retrieves the EndpointSlice from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.EndpointSlice, error)
	EndpointSliceNamespaceListerExpansion
}

// endpointSliceNamespaceLister implements the EndpointSliceNamespaceLister
// interface.
type endpointSliceNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  EndpointSlice from the indexer for a given workspace, namespace and name.
func (s endpointSliceNamespaceLister) Get(name string) (*v1.EndpointSlice, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("endpointslice"), name)
	}
	return obj.(*v1.EndpointSlice), nil
}

// List lists all EndpointSlices in the indexer for a given workspace, namespace and name.
func (s endpointSliceNamespaceLister) List(selector labels.Selector) (ret []*v1.EndpointSlice, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.EndpointSlice))
	})
	return ret, err
}
