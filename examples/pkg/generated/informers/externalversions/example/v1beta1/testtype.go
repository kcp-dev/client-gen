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

// Code generated by informer-gen-v0.32. DO NOT EDIT.

package v1beta1

import (
	context "context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	apisexamplev1beta1 "acme.corp/pkg/apis/example/v1beta1"
	versioned "acme.corp/pkg/generated/clientset/versioned"
	internalinterfaces "acme.corp/pkg/generated/informers/externalversions/internalinterfaces"
	examplev1beta1 "acme.corp/pkg/generated/listers/example/v1beta1"
)

// TestTypeInformer provides access to a shared informer and lister for
// TestTypes.
type TestTypeInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() examplev1beta1.TestTypeLister
}

type testTypeInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTestTypeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTestTypeInformer constructs a new informer for TestType type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTestTypeInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1beta1().TestTypes(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ExampleV1beta1().TestTypes(namespace).Watch(context.TODO(), options)
			},
		},
		&apisexamplev1beta1.TestType{},
		resyncPeriod,
		indexers,
	)
}

func (f *testTypeInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTestTypeInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *testTypeInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&apisexamplev1beta1.TestType{}, f.defaultInformer)
}

func (f *testTypeInformer) Lister() examplev1beta1.TestTypeLister {
	return examplev1beta1.NewTestTypeLister(f.Informer().GetIndexer())
}