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
	v1beta1 "k8s.io/api/apps/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// DeploymentLister helps list Deployments.
// All objects returned here must be treated as read-only.
type DeploymentLister interface {
	// List lists all Deployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Deployment, err error)
	// Deployments returns an object that can list and get Deployments.
	Deployments(namespace string) DeploymentNamespaceLister
	DeploymentListerExpansion
}

// DeploymentClusterLister helps list Deployments.
// All objects returned here must be treated as read-only.
type DeploymentClusterLister interface {
	// List lists all Deployments in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Deployment, err error)
	DeploymentClusterListerExpansion
}

// deploymentLister implements the DeploymentLister interface.
type deploymentLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// deploymentLister implements the DeploymentClusterLister interface.
type deploymentClusterLister struct {
	indexer cache.Indexer
}

// List lists all Deployments in the indexer.
func (s *deploymentLister) List(selector labels.Selector) (ret []*v1beta1.Deployment, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.Deployment))
	})
	return ret, err
}

// Get retrieves the  Deployment from the indexer for a given workspace, namespace and name.
func (s deploymentLister) Get(name string) (*v1beta1.Deployment, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("deployment"), name)
	}
	return obj.(*v1beta1.Deployment), nil
}

// NewDeploymentClusterLister returns a new DeploymentClusterLister.
func NewDeploymentClusterLister(indexer cache.Indexer) DeploymentClusterLister {
	return &deploymentClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get Deployment.
func (s *deploymentClusterLister) Cluster(clusterName logicalcluster.Name) DeploymentLister {
	return &deploymentLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all Deployments in the indexer.
func (s *deploymentClusterLister) List(selector labels.Selector) (ret []*v1beta1.Deployment, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.Deployment))
	})
	return ret, err
}

// Deployments returns an object that can list and get Deployments.
func (s *deploymentLister) Deployments(namespace string) DeploymentNamespaceLister {
	return deploymentNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// DeploymentNamespaceLister helps list and get Deployments.
// All objects returned here must be treated as read-only.
type DeploymentNamespaceLister interface {
	// List lists all Deployments in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.Deployment, err error)
	// Get retrieves the Deployment from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.Deployment, error)
	DeploymentNamespaceListerExpansion
}

// deploymentNamespaceLister implements the DeploymentNamespaceLister
// interface.
type deploymentNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  Deployment from the indexer for a given workspace, namespace and name.
func (s deploymentNamespaceLister) Get(name string) (*v1beta1.Deployment, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("deployment"), name)
	}
	return obj.(*v1beta1.Deployment), nil
}

// List lists all Deployments in the indexer for a given workspace, namespace and name.
func (s deploymentNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.Deployment, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.Deployment))
	})
	return ret, err
}
