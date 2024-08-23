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
	v1beta1 "k8s.io/api/storage/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StorageClassLister helps list StorageClasses.
// All objects returned here must be treated as read-only.
type StorageClassLister interface {
	// List lists all StorageClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.StorageClass, err error)
	// Get retrieves the StorageClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.StorageClass, error)
	StorageClassListerExpansion
}

// StorageClassClusterLister helps list StorageClasses.
// All objects returned here must be treated as read-only.
type StorageClassClusterLister interface {
	// List lists all StorageClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.StorageClass, err error)
	StorageClassClusterListerExpansion
}

// storageClassLister implements the StorageClassLister interface.
type storageClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// storageClassLister implements the StorageClassClusterLister interface.
type storageClassClusterLister struct {
	indexer cache.Indexer
}

// List lists all StorageClasses in the indexer.
func (s *storageClassLister) List(selector labels.Selector) (ret []*v1beta1.StorageClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.StorageClass))
	})
	return ret, err
}

// Get retrieves the  StorageClass from the indexer for a given workspace, namespace and name.
func (s storageClassLister) Get(name string) (*v1beta1.StorageClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("storageclass"), name)
	}
	return obj.(*v1beta1.StorageClass), nil
}

// NewStorageClassClusterLister returns a new StorageClassClusterLister.
func NewStorageClassClusterLister(indexer cache.Indexer) StorageClassClusterLister {
	return &storageClassClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get StorageClass.
func (s *storageClassClusterLister) Cluster(clusterName logicalcluster.Name) StorageClassLister {
	return &storageClassLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all StorageClasses in the indexer.
func (s *storageClassClusterLister) List(selector labels.Selector) (ret []*v1beta1.StorageClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.StorageClass))
	})
	return ret, err
}
