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

package v1beta1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1beta1 "k8s.io/api/rbac/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RoleBindingLister helps list RoleBindings.
// All objects returned here must be treated as read-only.
type RoleBindingLister interface {
	// List lists all RoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error)
	// RoleBindings returns an object that can list and get RoleBindings.
	RoleBindings(namespace string) RoleBindingNamespaceLister
	RoleBindingListerExpansion
}

// RoleBindingClusterLister helps list RoleBindings.
// All objects returned here must be treated as read-only.
type RoleBindingClusterLister interface {
	// List lists all RoleBindings in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error)
	RoleBindingClusterListerExpansion
}

// roleBindingLister implements the RoleBindingLister interface.
type roleBindingLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// roleBindingLister implements the RoleBindingClusterLister interface.
type roleBindingClusterLister struct {
	indexer cache.Indexer
}

// List lists all RoleBindings in the indexer.
func (s *roleBindingLister) List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.RoleBinding))
	})
	return ret, err
}

// Get retrieves the  RoleBinding from the indexer for a given workspace, namespace and name.
func (s roleBindingLister) Get(name string) (*v1beta1.RoleBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("rolebinding"), name)
	}
	return obj.(*v1beta1.RoleBinding), nil
}

// NewRoleBindingClusterLister returns a new RoleBindingClusterLister.
func NewRoleBindingClusterLister(indexer cache.Indexer) RoleBindingClusterLister {
	return &roleBindingClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get RoleBinding.
func (s *roleBindingClusterLister) Cluster(clusterName logicalcluster.Name) RoleBindingLister {
	return &roleBindingLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all RoleBindings in the indexer.
func (s *roleBindingClusterLister) List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.RoleBinding))
	})
	return ret, err
}

// RoleBindings returns an object that can list and get RoleBindings.
func (s *roleBindingLister) RoleBindings(namespace string) RoleBindingNamespaceLister {
	return roleBindingNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// RoleBindingNamespaceLister helps list and get RoleBindings.
// All objects returned here must be treated as read-only.
type RoleBindingNamespaceLister interface {
	// List lists all RoleBindings in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error)
	// Get retrieves the RoleBinding from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.RoleBinding, error)
	RoleBindingNamespaceListerExpansion
}

// roleBindingNamespaceLister implements the RoleBindingNamespaceLister
// interface.
type roleBindingNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  RoleBinding from the indexer for a given workspace, namespace and name.
func (s roleBindingNamespaceLister) Get(name string) (*v1beta1.RoleBinding, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("rolebinding"), name)
	}
	return obj.(*v1beta1.RoleBinding), nil
}

// List lists all RoleBindings in the indexer for a given workspace, namespace and name.
func (s roleBindingNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.RoleBinding, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.RoleBinding))
	})
	return ret, err
}