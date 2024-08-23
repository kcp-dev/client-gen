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
	v1beta1 "k8s.io/api/batch/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CronJobLister helps list CronJobs.
// All objects returned here must be treated as read-only.
type CronJobLister interface {
	// List lists all CronJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CronJob, err error)
	// CronJobs returns an object that can list and get CronJobs.
	CronJobs(namespace string) CronJobNamespaceLister
	CronJobListerExpansion
}

// CronJobClusterLister helps list CronJobs.
// All objects returned here must be treated as read-only.
type CronJobClusterLister interface {
	// List lists all CronJobs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CronJob, err error)
	CronJobClusterListerExpansion
}

// cronJobLister implements the CronJobLister interface.
type cronJobLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// cronJobLister implements the CronJobClusterLister interface.
type cronJobClusterLister struct {
	indexer cache.Indexer
}

// List lists all CronJobs in the indexer.
func (s *cronJobLister) List(selector labels.Selector) (ret []*v1beta1.CronJob, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.CronJob))
	})
	return ret, err
}

// Get retrieves the  CronJob from the indexer for a given workspace, namespace and name.
func (s cronJobLister) Get(name string) (*v1beta1.CronJob, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("cronjob"), name)
	}
	return obj.(*v1beta1.CronJob), nil
}

// NewCronJobClusterLister returns a new CronJobClusterLister.
func NewCronJobClusterLister(indexer cache.Indexer) CronJobClusterLister {
	return &cronJobClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get CronJob.
func (s *cronJobClusterLister) Cluster(clusterName logicalcluster.Name) CronJobLister {
	return &cronJobLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all CronJobs in the indexer.
func (s *cronJobClusterLister) List(selector labels.Selector) (ret []*v1beta1.CronJob, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.CronJob))
	})
	return ret, err
}

// CronJobs returns an object that can list and get CronJobs.
func (s *cronJobLister) CronJobs(namespace string) CronJobNamespaceLister {
	return cronJobNamespaceLister{indexer: s.indexer, clusterName: s.clusterName, namespace: namespace}
}

// CronJobNamespaceLister helps list and get CronJobs.
// All objects returned here must be treated as read-only.
type CronJobNamespaceLister interface {
	// List lists all CronJobs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.CronJob, err error)
	// Get retrieves the CronJob from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.CronJob, error)
	CronJobNamespaceListerExpansion
}

// cronJobNamespaceLister implements the CronJobNamespaceLister
// interface.
type cronJobNamespaceLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
	namespace   string
}

// Get retrieves the  CronJob from the indexer for a given workspace, namespace and name.
func (s cronJobNamespaceLister) Get(name string) (*v1beta1.CronJob, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("cronjob"), name)
	}
	return obj.(*v1beta1.CronJob), nil
}

// List lists all CronJobs in the indexer for a given workspace, namespace and name.
func (s cronJobNamespaceLister) List(selector labels.Selector) (ret []*v1beta1.CronJob, err error) {
	err = kcpcache.ListAllByClusterAndNamespace(s.indexer, s.clusterName, s.namespace, selector, func(i interface{}) {
		ret = append(ret, i.(*v1beta1.CronJob))
	})
	return ret, err
}
