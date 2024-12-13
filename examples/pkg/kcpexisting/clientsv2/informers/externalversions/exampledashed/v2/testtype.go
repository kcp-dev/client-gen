//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The KCP Authors.

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

// Code generated by kcp code-generator. DO NOT EDIT.

package v2

import (
	"context"
	"time"

	kcpcache "github.com/kcp-dev/apimachinery/v2/pkg/cache"
	kcpinformers "github.com/kcp-dev/apimachinery/v2/third_party/informers"
	"github.com/kcp-dev/logicalcluster/v3"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"

	exampledashedv2 "acme.corp/pkg/apisv2/exampledashed/v2"
	upstreamexampledashedv2informers "acme.corp/pkg/generated/informersv2/externalversions/exampledashed/v2"
	upstreamexampledashedv2listers "acme.corp/pkg/generated/listersv2/exampledashed/v2"
	clientset "acme.corp/pkg/kcpexisting/clientsv2/exampledashed/versioned"
	"acme.corp/pkg/kcpexisting/clientsv2/informers/externalversions/internalinterfaces"
	exampledashedv2listers "acme.corp/pkg/kcpexisting/clientsv2/listers/exampledashed/v2"
)

// TestTypeClusterInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeClusterInformer interface {
	Cluster(logicalcluster.Name) upstreamexampledashedv2informers.TestTypeInformer
	Informer() kcpcache.ScopeableSharedIndexInformer
	Lister() exampledashedv2listers.TestTypeClusterLister
}

type testTypeClusterInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeClusterInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeClusterInformer(client clientset.ClusterInterface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) kcpcache.ScopeableSharedIndexInformer {
	return kcpinformers.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleDashedV2().TestTypes().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleDashedV2().TestTypes().Watch(context.TODO(), options)
			},
		},
		&exampledashedv2.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeClusterInformer) defaultInformer(client clientset.ClusterInterface, resyncPeriod time.Duration) kcpcache.ScopeableSharedIndexInformer {
	return NewFilteredTestTypeClusterInformer(client, resyncPeriod, cache.Indexers{
		kcpcache.ClusterIndexName:             kcpcache.ClusterIndexFunc,
		kcpcache.ClusterAndNamespaceIndexName: kcpcache.ClusterAndNamespaceIndexFunc},
		f.tweakListOptions,
	)
}

func (f *testTypeClusterInformer) Informer() kcpcache.ScopeableSharedIndexInformer {
	return f.factory.InformerFor(&exampledashedv2.TestType{}, f.defaultInformer)
}

func (f *testTypeClusterInformer) Lister() exampledashedv2listers.TestTypeClusterLister {
	return exampledashedv2listers.NewTestTypeClusterLister(f.Informer().GetIndexer())
}

func (f *testTypeClusterInformer) Cluster(clusterName logicalcluster.Name) upstreamexampledashedv2informers.TestTypeInformer {
	return &testTypeInformer{
		informer: f.Informer().Cluster(clusterName),
		lister:   f.Lister().Cluster(clusterName),
	}
}

type testTypeInformer struct {
	informer cache.SharedIndexInformer
	lister   upstreamexampledashedv2listers.TestTypeLister
}

func (f *testTypeInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

func (f *testTypeInformer) Lister() upstreamexampledashedv2listers.TestTypeLister {
	return f.lister
}
