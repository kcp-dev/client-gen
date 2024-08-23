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

package v2beta2

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v2beta2 "k8s.io/api/autoscaling/v2beta2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HorizontalPodAutoscalerLister helps list HorizontalPodAutoscalers.
// All objects returned here must be treated as read-only.
type HorizontalPodAutoscalerLister interface {
	// List lists all HorizontalPodAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error)
	// HorizontalPodAutoscalers returns an object that can list and get HorizontalPodAutoscalers.
	HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerNamespaceLister
	HorizontalPodAutoscalerListerExpansion
}

// HorizontalPodAutoscalerClusterLister helps list HorizontalPodAutoscalers.
// All objects returned here must be treated as read-only.
type HorizontalPodAutoscalerClusterLister interface {
	// List lists all HorizontalPodAutoscalers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error)
	HorizontalPodAutoscalerClusterListerExpansion
}

// horizontalPodAutoscalerLister implements the HorizontalPodAutoscalerLister interface.
type horizontalPodAutoscalerLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// horizontalPodAutoscalerLister implements the HorizontalPodAutoscalerClusterLister interface.
type horizontalPodAutoscalerClusterLister struct {
	indexer cache.Indexer
}

// List lists all HorizontalPodAutoscalers in the indexer.
func (s *horizontalPodAutoscalerLister) List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v2beta2.HorizontalPodAutoscaler))
	})
	return ret, err
}

// Get retrieves the  HorizontalPodAutoscaler from the indexer for a given workspace, namespace and name.
func (s horizontalPodAutoscalerLister) Get(name string) (*v2beta2.HorizontalPodAutoscaler, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta2.Resource("horizontalpodautoscaler"), name)
	}
	return obj.(*v2beta2.HorizontalPodAutoscaler), nil
}

// NewHorizontalPodAutoscalerClusterLister returns a new HorizontalPodAutoscalerClusterLister.
func NewHorizontalPodAutoscalerClusterLister(indexer cache.Indexer) HorizontalPodAutoscalerClusterLister {
	return &horizontalPodAutoscalerClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get HorizontalPodAutoscaler.
func (s *horizontalPodAutoscalerClusterLister) Cluster(clusterName logicalcluster.Name) HorizontalPodAutoscalerLister {
	return &horizontalPodAutoscalerLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all HorizontalPodAutoscalers in the indexer.
func (s *horizontalPodAutoscalerClusterLister) List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2beta2.HorizontalPodAutoscaler))
	})
	return ret, err
}

// HorizontalPodAutoscalers returns an object that can list and get HorizontalPodAutoscalers.
func (s *horizontalPodAutoscalerLister) HorizontalPodAutoscalers(namespace string) HorizontalPodAutoscalerNamespaceLister {
	return horizontalPodAutoscalerNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// HorizontalPodAutoscalerNamespaceLister helps list and get HorizontalPodAutoscalers.
// All objects returned here must be treated as read-only.
type HorizontalPodAutoscalerNamespaceLister interface {
	// List lists all HorizontalPodAutoscalers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error)
	// Get retrieves the HorizontalPodAutoscaler from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2beta2.HorizontalPodAutoscaler, error)
	HorizontalPodAutoscalerNamespaceListerExpansion
}

// horizontalPodAutoscalerNamespaceLister implements the HorizontalPodAutoscalerNamespaceLister
// interface.
type horizontalPodAutoscalerNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  HorizontalPodAutoscaler from the indexer for a given workspace, namespace and name.
func (s horizontalPodAutoscalerNamespaceLister) Get(name string) (*v2beta2.HorizontalPodAutoscaler, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2beta2.Resource("horizontalpodautoscaler"), name)
	}
	return obj.(*v2beta2.HorizontalPodAutoscaler), nil
}

// List lists all HorizontalPodAutoscalers in the indexer for a given workspace, namespace and name.
func (s horizontalPodAutoscalerNamespaceLister) List(selector labels.Selector) (ret []*v2beta2.HorizontalPodAutoscaler, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v2beta2.HorizontalPodAutoscaler))
	})
	return ret, err
}