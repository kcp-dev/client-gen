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
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ReplicaSetLister helps list ReplicaSets.
// All objects returned here must be treated as read-only.
type ReplicaSetLister interface {
	// List lists all ReplicaSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ReplicaSet, err error)
	// ReplicaSets returns an object that can list and get ReplicaSets.
	ReplicaSets(namespace string) ReplicaSetNamespaceLister
	ReplicaSetListerExpansion
}

// ReplicaSetClusterLister helps list ReplicaSets.
// All objects returned here must be treated as read-only.
type ReplicaSetClusterLister interface {
	// List lists all ReplicaSets in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ReplicaSet, err error)
	ReplicaSetClusterListerExpansion
}

// replicaSetLister implements the ReplicaSetLister interface.
type replicaSetLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// replicaSetLister implements the ReplicaSetClusterLister interface.
type replicaSetClusterLister struct {
	indexer cache.Indexer
}

// List lists all ReplicaSets in the indexer.
func (s *replicaSetLister) List(selector labels.Selector) (ret []*v1.ReplicaSet, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ReplicaSet))
	})
	return ret, err
}

// Get retrieves the  ReplicaSet from the indexer for a given workspace, namespace and name.
func (s replicaSetLister) Get(name string) (*v1.ReplicaSet, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("replicaset"), name)
	}
	return obj.(*v1.ReplicaSet), nil
}

// NewReplicaSetClusterLister returns a new ReplicaSetClusterLister.
func NewReplicaSetClusterLister(indexer cache.Indexer) ReplicaSetClusterLister {
	return &replicaSetClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get ReplicaSet.
func (s *replicaSetClusterLister) Cluster(clusterName logicalcluster.Name) ReplicaSetLister {
	return &replicaSetLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all ReplicaSets in the indexer.
func (s *replicaSetClusterLister) List(selector labels.Selector) (ret []*v1.ReplicaSet, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ReplicaSet))
	})
	return ret, err
}

// ReplicaSets returns an object that can list and get ReplicaSets.
func (s *replicaSetLister) ReplicaSets(namespace string) ReplicaSetNamespaceLister {
	return replicaSetNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// ReplicaSetNamespaceLister helps list and get ReplicaSets.
// All objects returned here must be treated as read-only.
type ReplicaSetNamespaceLister interface {
	// List lists all ReplicaSets in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.ReplicaSet, err error)
	// Get retrieves the ReplicaSet from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.ReplicaSet, error)
	ReplicaSetNamespaceListerExpansion
}

// replicaSetNamespaceLister implements the ReplicaSetNamespaceLister
// interface.
type replicaSetNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  ReplicaSet from the indexer for a given workspace, namespace and name.
func (s replicaSetNamespaceLister) Get(name string) (*v1.ReplicaSet, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("replicaset"), name)
	}
	return obj.(*v1.ReplicaSet), nil
}

// List lists all ReplicaSets in the indexer for a given workspace, namespace and name.
func (s replicaSetNamespaceLister) List(selector labels.Selector) (ret []*v1.ReplicaSet, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.ReplicaSet))
	})
	return ret, err
}