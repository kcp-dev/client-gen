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
	v1 "k8s.io/api/storage/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CSIDriverLister helps list CSIDrivers.
// All objects returned here must be treated as read-only.
type CSIDriverLister interface {
	// List lists all CSIDrivers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CSIDriver, err error)
	// Get retrieves the CSIDriver from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CSIDriver, error)
	CSIDriverListerExpansion
}

// CSIDriverClusterLister helps list CSIDrivers.
// All objects returned here must be treated as read-only.
type CSIDriverClusterLister interface {
	// List lists all CSIDrivers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CSIDriver, err error)
	CSIDriverClusterListerExpansion
}

// cSIDriverLister implements the CSIDriverLister interface.
type cSIDriverLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// cSIDriverLister implements the CSIDriverClusterLister interface.
type cSIDriverClusterLister struct {
	indexer cache.Indexer
}

// List lists all CSIDrivers in the indexer.
func (s *cSIDriverLister) List(selector labels.Selector) (ret []*v1.CSIDriver, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.CSIDriver))
	})
	return ret, err
}

// Get retrieves the  CSIDriver from the indexer for a given workspace, namespace and name.
func (s cSIDriverLister) Get(name string) (*v1.CSIDriver, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("csidriver"), name)
	}
	return obj.(*v1.CSIDriver), nil
}

// NewCSIDriverClusterLister returns a new CSIDriverClusterLister.
func NewCSIDriverClusterLister(indexer cache.Indexer) CSIDriverClusterLister {
	return &cSIDriverClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get CSIDriver.
func (s *cSIDriverClusterLister) Cluster(clusterName logicalcluster.Name) CSIDriverLister {
	return &cSIDriverLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all CSIDrivers in the indexer.
func (s *cSIDriverClusterLister) List(selector labels.Selector) (ret []*v1.CSIDriver, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CSIDriver))
	})
	return ret, err
}
