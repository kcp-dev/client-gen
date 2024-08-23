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
	v1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// NetworkPolicyLister helps list NetworkPolicies.
// All objects returned here must be treated as read-only.
type NetworkPolicyLister interface {
	// List lists all NetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error)
	// NetworkPolicies returns an object that can list and get NetworkPolicies.
	NetworkPolicies(namespace string) NetworkPolicyNamespaceLister
	NetworkPolicyListerExpansion
}

// NetworkPolicyClusterLister helps list NetworkPolicies.
// All objects returned here must be treated as read-only.
type NetworkPolicyClusterLister interface {
	// List lists all NetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error)
	NetworkPolicyClusterListerExpansion
}

// networkPolicyLister implements the NetworkPolicyLister interface.
type networkPolicyLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// networkPolicyLister implements the NetworkPolicyClusterLister interface.
type networkPolicyClusterLister struct {
	indexer cache.Indexer
}

// List lists all NetworkPolicies in the indexer.
func (s *networkPolicyLister) List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.NetworkPolicy))
	})
	return ret, err
}

// Get retrieves the  NetworkPolicy from the indexer for a given workspace, namespace and name.
func (s networkPolicyLister) Get(name string) (*v1.NetworkPolicy, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("networkpolicy"), name)
	}
	return obj.(*v1.NetworkPolicy), nil
}

// NewNetworkPolicyClusterLister returns a new NetworkPolicyClusterLister.
func NewNetworkPolicyClusterLister(indexer cache.Indexer) NetworkPolicyClusterLister {
	return &networkPolicyClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get NetworkPolicy.
func (s *networkPolicyClusterLister) Cluster(clusterName logicalcluster.Name) NetworkPolicyLister {
	return &networkPolicyLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all NetworkPolicies in the indexer.
func (s *networkPolicyClusterLister) List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkPolicy))
	})
	return ret, err
}

// NetworkPolicies returns an object that can list and get NetworkPolicies.
func (s *networkPolicyLister) NetworkPolicies(namespace string) NetworkPolicyNamespaceLister {
	return networkPolicyNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// NetworkPolicyNamespaceLister helps list and get NetworkPolicies.
// All objects returned here must be treated as read-only.
type NetworkPolicyNamespaceLister interface {
	// List lists all NetworkPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error)
	// Get retrieves the NetworkPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NetworkPolicy, error)
	NetworkPolicyNamespaceListerExpansion
}

// networkPolicyNamespaceLister implements the NetworkPolicyNamespaceLister
// interface.
type networkPolicyNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  NetworkPolicy from the indexer for a given workspace, namespace and name.
func (s networkPolicyNamespaceLister) Get(name string) (*v1.NetworkPolicy, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("networkpolicy"), name)
	}
	return obj.(*v1.NetworkPolicy), nil
}

// List lists all NetworkPolicies in the indexer for a given workspace, namespace and name.
func (s networkPolicyNamespaceLister) List(selector labels.Selector) (ret []*v1.NetworkPolicy, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.NetworkPolicy))
	})
	return ret, err
}
