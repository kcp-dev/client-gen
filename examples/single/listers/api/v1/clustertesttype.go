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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/code-generator/examples/single/api/v1"
)

// ClusterTestTypeLister helps list ClusterTestTypes.
// All objects returned here must be treated as read-only.
type ClusterTestTypeLister interface {
	// List lists all ClusterTestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTestType, err error)
	// Get retrieves the ClusterTestType from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ClusterTestType, error)
	ClusterTestTypeListerExpansion
}

// ClusterTestTypeClusterLister helps list ClusterTestTypes.
// All objects returned here must be treated as read-only.
type ClusterTestTypeClusterLister interface {
	// List lists all ClusterTestTypes in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ClusterTestType, err error)
	ClusterTestTypeClusterListerExpansion
}

// clusterTestTypeLister implements the ClusterTestTypeLister interface.
type clusterTestTypeLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// clusterTestTypeLister implements the ClusterTestTypeClusterLister interface.
type clusterTestTypeClusterLister struct {
	indexer cache.Indexer
}

// List lists all ClusterTestTypes in the indexer.
func (s *clusterTestTypeLister) List(selector labels.Selector) (ret []*v1.ClusterTestType, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ClusterTestType))
	})
	return ret, err
}

// Get retrieves the  ClusterTestType from the indexer for a given workspace, namespace and name.
func (s clusterTestTypeLister) Get(name string) (*v1.ClusterTestType, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("clustertesttype"), name)
	}
	return obj.(*v1.ClusterTestType), nil
}

// NewClusterTestTypeClusterLister returns a new ClusterTestTypeClusterLister.
func NewClusterTestTypeClusterLister(indexer cache.Indexer) ClusterTestTypeClusterLister {
	return &clusterTestTypeClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get ClusterTestType.
func (s *clusterTestTypeClusterLister) Cluster(clusterName logicalcluster.Name) ClusterTestTypeLister {
	return &clusterTestTypeLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ClusterTestTypes in the indexer.
func (s *clusterTestTypeClusterLister) List(selector labels.Selector) (ret []*v1.ClusterTestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ClusterTestType))
	})
	return ret, err
}