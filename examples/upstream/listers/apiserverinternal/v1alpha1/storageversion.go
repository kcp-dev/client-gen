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

package v1alpha1

import (
	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	"github.com/kcp-dev/logicalcluster/v3"
	v1alpha1 "k8s.io/api/apiserverinternal/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageVersionLister helps list StorageVersions.
// All objects returned here must be treated as read-only.
type StorageVersionLister interface {
	// List lists all StorageVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageVersion, err error)
	// Get retrieves the StorageVersion from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.StorageVersion, error)
	StorageVersionListerExpansion
}

// StorageVersionClusterLister helps list StorageVersions.
// All objects returned here must be treated as read-only.
type StorageVersionClusterLister interface {
	// List lists all StorageVersions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.StorageVersion, err error)
	StorageVersionClusterListerExpansion
}

// storageVersionLister implements the StorageVersionLister interface.
type storageVersionLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// storageVersionLister implements the StorageVersionClusterLister interface.
type storageVersionClusterLister struct {
	indexer cache.Indexer
}

// List lists all StorageVersions in the indexer.
func (s *storageVersionLister) List(selector labels.Selector) (ret []*v1alpha1.StorageVersion, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha1.StorageVersion))
	})
	return ret, err
}

// Get retrieves the  StorageVersion from the indexer for a given workspace, namespace and name.
func (s storageVersionLister) Get(name string) (*v1alpha1.StorageVersion, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("storageversion"), name)
	}
	return obj.(*v1alpha1.StorageVersion), nil
}

// NewStorageVersionClusterLister returns a new StorageVersionClusterLister.
func NewStorageVersionClusterLister(indexer cache.Indexer) StorageVersionClusterLister {
	return &storageVersionClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get StorageVersion.
func (s *storageVersionClusterLister) Cluster(clusterName logicalcluster.Name) StorageVersionLister {
	return &storageVersionLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all StorageVersions in the indexer.
func (s *storageVersionClusterLister) List(selector labels.Selector) (ret []*v1alpha1.StorageVersion, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.StorageVersion))
	})
	return ret, err
}
