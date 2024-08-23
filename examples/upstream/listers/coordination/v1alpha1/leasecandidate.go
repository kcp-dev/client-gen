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
	v1alpha1 "k8s.io/api/coordination/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// LeaseCandidateLister helps list LeaseCandidates.
// All objects returned here must be treated as read-only.
type LeaseCandidateLister interface {
	// List lists all LeaseCandidates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error)
	// LeaseCandidates returns an object that can list and get LeaseCandidates.
	LeaseCandidates(namespace string) LeaseCandidateNamespaceLister
	LeaseCandidateListerExpansion
}

// LeaseCandidateClusterLister helps list LeaseCandidates.
// All objects returned here must be treated as read-only.
type LeaseCandidateClusterLister interface {
	// List lists all LeaseCandidates in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error)
	LeaseCandidateClusterListerExpansion
}

// leaseCandidateLister implements the LeaseCandidateLister interface.
type leaseCandidateLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// leaseCandidateLister implements the LeaseCandidateClusterLister interface.
type leaseCandidateClusterLister struct {
	indexer cache.Indexer
}

// List lists all LeaseCandidates in the indexer.
func (s *leaseCandidateLister) List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha1.LeaseCandidate))
	})
	return ret, err
}

// Get retrieves the  LeaseCandidate from the indexer for a given workspace, namespace and name.
func (s leaseCandidateLister) Get(name string) (*v1alpha1.LeaseCandidate, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("leasecandidate"), name)
	}
	return obj.(*v1alpha1.LeaseCandidate), nil
}

// NewLeaseCandidateClusterLister returns a new LeaseCandidateClusterLister.
func NewLeaseCandidateClusterLister(indexer cache.Indexer) LeaseCandidateClusterLister {
	return &leaseCandidateClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get LeaseCandidate.
func (s *leaseCandidateClusterLister) Cluster(clusterName logicalcluster.Name) LeaseCandidateLister {
	return &leaseCandidateLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all LeaseCandidates in the indexer.
func (s *leaseCandidateClusterLister) List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.LeaseCandidate))
	})
	return ret, err
}

// LeaseCandidates returns an object that can list and get LeaseCandidates.
func (s *leaseCandidateLister) LeaseCandidates(namespace string) LeaseCandidateNamespaceLister {
	return leaseCandidateNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// LeaseCandidateNamespaceLister helps list and get LeaseCandidates.
// All objects returned here must be treated as read-only.
type LeaseCandidateNamespaceLister interface {
	// List lists all LeaseCandidates in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error)
	// Get retrieves the LeaseCandidate from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.LeaseCandidate, error)
	LeaseCandidateNamespaceListerExpansion
}

// leaseCandidateNamespaceLister implements the LeaseCandidateNamespaceLister
// interface.
type leaseCandidateNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  LeaseCandidate from the indexer for a given workspace, namespace and name.
func (s leaseCandidateNamespaceLister) Get(name string) (*v1alpha1.LeaseCandidate, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("leasecandidate"), name)
	}
	return obj.(*v1alpha1.LeaseCandidate), nil
}

// List lists all LeaseCandidates in the indexer for a given workspace, namespace and name.
func (s leaseCandidateNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.LeaseCandidate, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1alpha1.LeaseCandidate))
	})
	return ret, err
}
