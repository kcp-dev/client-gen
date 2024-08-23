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
	v1 "k8s.io/api/flowcontrol/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// FlowSchemaLister helps list FlowSchemas.
// All objects returned here must be treated as read-only.
type FlowSchemaLister interface {
	// List lists all FlowSchemas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.FlowSchema, err error)
	// Get retrieves the FlowSchema from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.FlowSchema, error)
	FlowSchemaListerExpansion
}

// FlowSchemaClusterLister helps list FlowSchemas.
// All objects returned here must be treated as read-only.
type FlowSchemaClusterLister interface {
	// List lists all FlowSchemas in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.FlowSchema, err error)
	FlowSchemaClusterListerExpansion
}

// flowSchemaLister implements the FlowSchemaLister interface.
type flowSchemaLister struct {
	indexer     cache.Indexer
	clusterName logicalcluster.Name
}

// flowSchemaLister implements the FlowSchemaClusterLister interface.
type flowSchemaClusterLister struct {
	indexer cache.Indexer
}

// List lists all FlowSchemas in the indexer.
func (s *flowSchemaLister) List(selector labels.Selector) (ret []*v1.FlowSchema, err error) {
	err = kcpcache.ListAllByCluster(s.indexer, s.clusterName, selector, func(i interface{}) {
		ret = append(ret, i.(*v1.FlowSchema))
	})
	return ret, err
}

// Get retrieves the  FlowSchema from the indexer for a given workspace, namespace and name.
func (s flowSchemaLister) Get(name string) (*v1.FlowSchema, error) {
	key := kcpcache.ToClusterAwareKey(s.clusterName.String(), "", name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("flowschema"), name)
	}
	return obj.(*v1.FlowSchema), nil
}

// NewFlowSchemaClusterLister returns a new FlowSchemaClusterLister.
func NewFlowSchemaClusterLister(indexer cache.Indexer) FlowSchemaClusterLister {
	return &flowSchemaClusterLister{indexer: indexer}
}

// Cluster scopes the lister to one workspace, allowing users to list and get FlowSchema.
func (s *flowSchemaClusterLister) Cluster(clusterName logicalcluster.Name) FlowSchemaLister {
	return &flowSchemaLister{indexer: s.indexer, clusterName: clusterName}
}

// List lists all FlowSchemas in the indexer.
func (s *flowSchemaClusterLister) List(selector labels.Selector) (ret []*v1.FlowSchema, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.FlowSchema))
	})
	return ret, err
}