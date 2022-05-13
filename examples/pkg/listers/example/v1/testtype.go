
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code auto-generated. DO NOT EDIT.

package v1



import (
	examplev1 "github.com/kcp-dev/code-generator/examples/pkg/apis/example/v1"
	"github.com/kcp-dev/logicalcluster"
	"k8s.io/client-go/tools/cache"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/api/errors"
	apimachinerycache "github.com/kcp-dev/apimachinery/pkg/cache"
)



// TestTypeLister helps list testType.
// All objects returned here must be treated as read-only.
type TestTypeClusterLister interface {
	// List lists all testType in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev1.TestType, err error)

	// Cluster returns an object that can list and get testType from the given logical cluster.
	Cluster(cluster logicalcluster.Name) TestTypeLister

	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// TestTypeListerExpansion
}

// testTypeClusterLister implements the TestTypeClusterLister interface.
type testTypeClusterLister struct {
	indexer cache.Indexer
}

// NewTestTypeClusterLister returns a new TestTypeClusterLister.
func NewTestTypeClusterLister(indexer cache.Indexer) TestTypeClusterLister {
	return &testTypeClusterLister{indexer: indexer}
}

// List lists all testType in the indexer.
func (s *testTypeClusterLister) List(selector labels.Selector) (ret []*examplev1.TestType, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*examplev1.TestType))
	})
	return ret, err
}

// Cluster returns an object that can list and get testType.
func (s *testTypeClusterLister) Cluster(cluster logicalcluster.Name) TestTypeLister {
	return &testTypeLister{indexer: s.indexer, cluster: cluster}
}

// TestTypeLister helps list testType.
// All objects returned here must be treated as read-only.
type TestTypeLister interface {
	// List lists all testType in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev1.TestType, err error)
	// TestTypes returns an object that can list and get testType.
	TestTypes(namespace string) TestTypeNamespaceLister
	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// TestTypeListerExpansion
}

// testTypeLister implements the TestTypeLister interface.
type testTypeLister struct {
	indexer cache.Indexer
	cluster logicalcluster.Name
}

// List lists all testType in the indexer.
func (s *testTypeLister) List(selector labels.Selector) (ret []*examplev1.TestType, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), "", "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*examplev1.TestType)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}

	return ret, err
}

// testType returns an object that can list and get testType.
func (s *testTypeLister) TestTypes(namespace string) TestTypeNamespaceLister {
	return testTypeNamespaceLister{indexer: s.indexer, cluster: s.cluster, namespace: namespace}
}

// TestTypeNamespaceLister helps list and get testType.
// All objects returned here must be treated as read-only.
type TestTypeNamespaceLister interface {
	// List lists all testType in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*examplev1.TestType, err error)
	// Get retrieves the TestType from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*examplev1.TestType, error)
	// Note(kcp): Workspace-capable Lister implementation doesn't support support expansions.
	// TestTypeNamespaceListerExpansion
}

// testTypeNamespaceLister implements the TestTypeNamespaceLister
// interface.
type testTypeNamespaceLister struct {
	indexer   cache.Indexer
	cluster   logicalcluster.Name
	namespace string
}

// List lists all testType in the indexer for a given namespace.
func (s testTypeNamespaceLister) List(selector labels.Selector) (ret []*examplev1.TestType, err error) {
	selectAll := selector == nil || selector.Empty()

	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), s.namespace, "")
	list, err := s.indexer.ByIndex(apimachinerycache.ClusterAndNamespaceIndexName, key)
	if err != nil {
		return nil, err
	}

	for i := range list {
		obj := list[i].(*examplev1.TestType)
		if selectAll {
			ret = append(ret, obj)
		} else {
			if selector.Matches(labels.Set(obj.GetLabels())) {
				ret = append(ret, obj)
			}
		}
	}
	return ret, err
}

// Get retrieves the TestType from the indexer for a given namespace and name.
func (s testTypeNamespaceLister) Get(name string) (*examplev1.TestType, error) {
	key := apimachinerycache.ToClusterAwareKey(s.cluster.String(), s.namespace, name)
	obj, exists, err := s.indexer.GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(examplev1.Resource("testType"), name)
	}
	return obj.(*examplev1.TestType), nil
}
