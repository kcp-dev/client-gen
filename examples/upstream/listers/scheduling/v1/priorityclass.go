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
	v1 "k8s.io/api/scheduling/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PriorityClassLister helps list PriorityClasses.
// All objects returned here must be treated as read-only.
type PriorityClassLister interface {
	// List lists all PriorityClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PriorityClass, err error)
	// Get retrieves the PriorityClass from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.PriorityClass, error)
	PriorityClassListerExpansion
}

// PriorityClassClusterLister helps list PriorityClasses.
// All objects returned here must be treated as read-only.
type PriorityClassClusterLister interface {
	// List lists all PriorityClasses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.PriorityClass, err error)
	PriorityClassClusterListerExpansion
}

// priorityClassLister implements the PriorityClassLister interface.
type priorityClassLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// priorityClassLister implements the PriorityClassClusterLister interface.
type priorityClassClusterLister struct {
	indexer cache.Indexer
}

// List lists all PriorityClasses in the indexer.
func (s *priorityClassLister) List(selector labels.Selector) (ret []*v1.PriorityClass, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.PriorityClass))
	})
	return ret, err
}

// Get retrieves the  PriorityClass from the indexer for a given workspace, namespace and name.
func (s priorityClassLister) Get(name string) (*v1.PriorityClass, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("priorityclass"), name)
	}
	return obj.(*v1.PriorityClass), nil
}

// NewPriorityClassClusterLister returns a new PriorityClassClusterLister.
func NewPriorityClassClusterLister(indexer cache.Indexer) PriorityClassClusterLister {
	return &priorityClassClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get PriorityClass.
func (s *priorityClassClusterLister) Cluster(clusterName logicalcluster.Name) PriorityClassLister {
	return &priorityClassLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all PriorityClasses in the indexer.
func (s *priorityClassClusterLister) List(selector labels.Selector) (ret []*v1.PriorityClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.PriorityClass))
	})
	return ret, err
}
