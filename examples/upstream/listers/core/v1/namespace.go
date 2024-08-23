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

// NamespaceLister helps list Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	// Get retrieves the Namespace from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.Namespace, error)
	NamespaceListerExpansion
}

// NamespaceClusterLister helps list Namespaces.
// All objects returned here must be treated as read-only.
type NamespaceClusterLister interface {
	// List lists all Namespaces in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.Namespace, err error)
	NamespaceClusterListerExpansion
}

// namespaceLister implements the NamespaceLister interface.
type namespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// namespaceLister implements the NamespaceClusterLister interface.
type namespaceClusterLister struct {
	indexer cache.Indexer
}

// List lists all Namespaces in the indexer.
func (s *namespaceLister) List(selector labels.Selector) (ret []*v1.Namespace, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.Namespace))
	})
	return ret, err
}

// Get retrieves the  Namespace from the indexer for a given workspace, namespace and name.
func (s namespaceLister) Get(name string) (*v1.Namespace, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("namespace"), name)
	}
	return obj.(*v1.Namespace), nil
}

// NewNamespaceClusterLister returns a new NamespaceClusterLister.
func NewNamespaceClusterLister(indexer cache.Indexer) NamespaceClusterLister {
	return &namespaceClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get Namespace.
func (s *namespaceClusterLister) Cluster(clusterName logicalcluster.Name) NamespaceLister {
	return &namespaceLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all Namespaces in the indexer.
func (s *namespaceClusterLister) List(selector labels.Selector) (ret []*v1.Namespace, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.Namespace))
	})
	return ret, err
}
