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
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ConfigMapLister helps list ConfigMaps.
// All objects returned here must be treated as read-only.
type ConfigMapLister interface {
	// List lists all ConfigMaps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConfigMap, err error)
	// ConfigMaps returns an object that can list and get ConfigMaps.
	ConfigMaps(namespace string) ConfigMapNamespaceLister
	ConfigMapListerExpansion
}

// ConfigMapClusterLister helps list ConfigMaps.
// All objects returned here must be treated as read-only.
type ConfigMapClusterLister interface {
	// List lists all ConfigMaps in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConfigMap, err error)
	ConfigMapClusterListerExpansion
}

// configMapLister implements the ConfigMapLister interface.
type configMapLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// configMapLister implements the ConfigMapClusterLister interface.
type configMapClusterLister struct {
	indexer cache.Indexer
}

// List lists all ConfigMaps in the indexer.
func (s *configMapLister) List(selector labels.Selector) (ret []*v1.ConfigMap, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ConfigMap))
	})
	return ret, err
}

// Get retrieves the  ConfigMap from the indexer for a given workspace, namespace and name.
func (s configMapLister) Get(name string) (*v1.ConfigMap, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("configmap"), name)
	}
	return obj.(*v1.ConfigMap), nil
}

// NewConfigMapClusterLister returns a new ConfigMapClusterLister.
func NewConfigMapClusterLister(indexer cache.Indexer) ConfigMapClusterLister {
	return &configMapClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get ConfigMap.
func (s *configMapClusterLister) Cluster(clusterName logicalcluster.Name) ConfigMapLister {
	return &configMapLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ConfigMaps in the indexer.
func (s *configMapClusterLister) List(selector labels.Selector) (ret []*v1.ConfigMap, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ConfigMap))
	})
	return ret, err
}

// ConfigMaps returns an object that can list and get ConfigMaps.
func (s *configMapLister) ConfigMaps(namespace string) ConfigMapNamespaceLister {
	return configMapNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// ConfigMapNamespaceLister helps list and get ConfigMaps.
// All objects returned here must be treated as read-only.
type ConfigMapNamespaceLister interface {
	// List lists all ConfigMaps in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ConfigMap, err error)
	// Get retrieves the ConfigMap from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ConfigMap, error)
	ConfigMapNamespaceListerExpansion
}

// configMapNamespaceLister implements the ConfigMapNamespaceLister
// interface.
type configMapNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  ConfigMap from the indexer for a given workspace, namespace and name.
func (s configMapNamespaceLister) Get(name string) (*v1.ConfigMap, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("configmap"), name)
	}
	return obj.(*v1.ConfigMap), nil
}

// List lists all ConfigMaps in the indexer for a given workspace, namespace and name.
func (s configMapNamespaceLister) List(selector labels.Selector) (ret []*v1.ConfigMap, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ConfigMap))
	})
	return ret, err
}