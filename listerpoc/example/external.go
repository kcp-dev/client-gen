package example

import (
	"fmt"
	"strings"

	"github.com/kcp-dev/apimachinery/pkg/logicalcluster"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/tools/cache"
)

// TODO: All code here needs to move to an external library

const (
	ClusterIndexName             = "cluster"
	ClusterAndNamespaceIndexName = "cluster-and-namespace"
)

func ClusterIndexFunc(obj interface{}) ([]string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return []string{""}, fmt.Errorf("object has no meta: %v", err)
	}
	// return []string{meta.GetZZZ_DeprecatedClusterName()}, nil
	index := []string{meta.GetZZZ_DeprecatedClusterName()}
	fmt.Printf("\t\tClusterIndexFunc -> %v\n", index)
	return index, nil
}

func ClusterAndNamespaceIndexFunc(obj interface{}) ([]string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return []string{""}, fmt.Errorf("object has no meta: %v", err)
	}
	// return []string{meta.GetZZZ_DeprecatedClusterName() + "/" + meta.GetNamespace()}, nil
	index := []string{meta.GetZZZ_DeprecatedClusterName() + "/" + meta.GetNamespace()}
	fmt.Printf("\t\tClusterAndNamespaceIndexFunc -> %v\n", index)
	return index, nil

}

func ClusterAwareKeyFunc(obj interface{}) (string, error) {
	meta, err := meta.Accessor(obj)
	if err != nil {
		return "", fmt.Errorf("object has no meta: %v", err)
	}
	clusterName := meta.GetZZZ_DeprecatedClusterName()
	namespace := meta.GetNamespace()
	name := meta.GetName()

	return strings.Join([]string{clusterName, namespace, name}, "/"), nil
}

type GenericLister interface {
	// List will return all objects across clusters
	List(selector labels.Selector) (ret []runtime.Object, err error)
	// Get will attempt to retrieve assuming that name==key
	Get(name string) (runtime.Object, error)
	// ByCluster will give you a GenericClusterLister for one namespace
	ByCluster(cluster logicalcluster.LogicalCluster) cache.GenericLister
}

func NewGenericLister(indexer cache.Indexer, resource schema.GroupResource) GenericLister {
	return &genericLister{indexer: indexer, resource: resource}
}

type genericLister struct {
	indexer  cache.Indexer
	resource schema.GroupResource
}

func (s *genericLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(runtime.Object))
	})
	return ret, err
}

func (s *genericLister) Get(name string) (runtime.Object, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(runtime.Object), nil
}

func (s *genericLister) ByCluster(cluster logicalcluster.LogicalCluster) cache.GenericLister {
	return &genericClusterLister{indexer: s.indexer, resource: s.resource, cluster: cluster}
}

type genericClusterLister struct {
	indexer  cache.Indexer
	cluster  logicalcluster.LogicalCluster
	resource schema.GroupResource
}

func (s *genericClusterLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	selectAll := selector.Empty()
	list, err := s.indexer.ByIndex(ClusterIndexName, s.cluster.String())
	if err != nil {
		return nil, err
	}

	if selector == nil {
		selector = labels.Everything()
	}
	for i := range list {
		item := list[i].(runtime.Object)
		if selectAll {
			ret = append(ret, item)
		} else {
			metadata, err := meta.Accessor(item)
			if err != nil {
				return nil, err
			}
			if selector.Matches(labels.Set(metadata.GetLabels())) {
				ret = append(ret, item)
			}
		}
	}

	return ret, err
}

func (s *genericClusterLister) Get(name string) (runtime.Object, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(runtime.Object), nil
}

func (s *genericClusterLister) ByNamespace(namespace string) cache.GenericNamespaceLister {
	return &genericNamespaceLister{indexer: s.indexer, namespace: namespace, resource: s.resource, cluster: s.cluster}
}

type genericNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.LogicalCluster
	namespace string
	resource  schema.GroupResource
}

func (s *genericNamespaceLister) List(selector labels.Selector) (ret []runtime.Object, err error) {
	selectAll := selector.Empty()
	list, err := s.indexer.Index(ClusterAndNamespaceIndexName, &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: s.cluster.String(),
		Namespace:                 s.namespace,
	})
	if err != nil {
		return nil, err
	}

	for i := range list {
		item := list[i].(runtime.Object)
		if selectAll {
			ret = append(ret, item)
		} else {
			metadata, err := meta.Accessor(item)
			if err != nil {
				return nil, err
			}
			if selector.Matches(labels.Set(metadata.GetLabels())) {
				ret = append(ret, item)
			}
		}
	}
	return ret, err
}

func (s *genericNamespaceLister) Get(name string) (runtime.Object, error) {
	meta := &metav1.ObjectMeta{
		ZZZ_DeprecatedClusterName: s.cluster.String(),
		Namespace:                 s.namespace,
		Name:                      name,
	}
	obj, exists, err := s.indexer.Get(meta)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(s.resource, name)
	}
	return obj.(runtime.Object), nil
}
