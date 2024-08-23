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

// PersistentVolumeClaimLister helps list PersistentVolumeClaims.
// All objects returned here must be treated as read-only.
type PersistentVolumeClaimLister interface {
	// List lists all PersistentVolumeClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	// PersistentVolumeClaims returns an object that can list and get PersistentVolumeClaims.
	PersistentVolumeClaims(namespace string) PersistentVolumeClaimNamespaceLister
	PersistentVolumeClaimListerExpansion
}

// PersistentVolumeClaimClusterLister helps list PersistentVolumeClaims.
// All objects returned here must be treated as read-only.
type PersistentVolumeClaimClusterLister interface {
	// List lists all PersistentVolumeClaims in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	PersistentVolumeClaimClusterListerExpansion
}

// persistentVolumeClaimLister implements the PersistentVolumeClaimLister interface.
type persistentVolumeClaimLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// persistentVolumeClaimLister implements the PersistentVolumeClaimClusterLister interface.
type persistentVolumeClaimClusterLister struct {
	indexer cache.Indexer
}

// List lists all PersistentVolumeClaims in the indexer.
func (s *persistentVolumeClaimLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.PersistentVolumeClaim))
	})
	return ret, err
}

// Get retrieves the  PersistentVolumeClaim from the indexer for a given workspace, namespace and name.
func (s persistentVolumeClaimLister) Get(name string) (*v1.PersistentVolumeClaim, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("persistentvolumeclaim"), name)
	}
	return obj.(*v1.PersistentVolumeClaim), nil
}

// NewPersistentVolumeClaimClusterLister returns a new PersistentVolumeClaimClusterLister.
func NewPersistentVolumeClaimClusterLister(indexer cache.Indexer) PersistentVolumeClaimClusterLister {
	return &persistentVolumeClaimClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get PersistentVolumeClaim.
func (s *persistentVolumeClaimClusterLister) Cluster(clusterName logicalcluster.Name) PersistentVolumeClaimLister {
	return &persistentVolumeClaimLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all PersistentVolumeClaims in the indexer.
func (s *persistentVolumeClaimClusterLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PersistentVolumeClaim))
	})
	return ret, err
}

// PersistentVolumeClaims returns an object that can list and get PersistentVolumeClaims.
func (s *persistentVolumeClaimLister) PersistentVolumeClaims(namespace string) PersistentVolumeClaimNamespaceLister {
	return persistentVolumeClaimNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// PersistentVolumeClaimNamespaceLister helps list and get PersistentVolumeClaims.
// All objects returned here must be treated as read-only.
type PersistentVolumeClaimNamespaceLister interface {
	// List lists all PersistentVolumeClaims in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error)
	// Get retrieves the PersistentVolumeClaim from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PersistentVolumeClaim, error)
	PersistentVolumeClaimNamespaceListerExpansion
}

// persistentVolumeClaimNamespaceLister implements the PersistentVolumeClaimNamespaceLister
// interface.
type persistentVolumeClaimNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  PersistentVolumeClaim from the indexer for a given workspace, namespace and name.
func (s persistentVolumeClaimNamespaceLister) Get(name string) (*v1.PersistentVolumeClaim, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("persistentvolumeclaim"), name)
	}
	return obj.(*v1.PersistentVolumeClaim), nil
}

// List lists all PersistentVolumeClaims in the indexer for a given workspace, namespace and name.
func (s persistentVolumeClaimNamespaceLister) List(selector labels.Selector) (ret []*v1.PersistentVolumeClaim, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.PersistentVolumeClaim))
	})
	return ret, err
}
