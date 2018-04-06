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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	tryapi_v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
	versioned "github.com/nightfury1204/test-apiserver/pkg/client/clientset/versioned"
	internalinterfaces "github.com/nightfury1204/test-apiserver/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/client/listers/tryapi/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PlayerInformer provides access to a shared informer and lister for
// Players.
type PlayerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.PlayerLister
}

type playerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPlayerInformer constructs a new informer for Player type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPlayerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPlayerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPlayerInformer constructs a new informer for Player type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPlayerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TryapiV1alpha1().Players(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TryapiV1alpha1().Players(namespace).Watch(options)
			},
		},
		&tryapi_v1alpha1.Player{},
		resyncPeriod,
		indexers,
	)
}

func (f *playerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPlayerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *playerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&tryapi_v1alpha1.Player{}, f.defaultInformer)
}

func (f *playerInformer) Lister() v1alpha1.PlayerLister {
	return v1alpha1.NewPlayerLister(f.Informer().GetIndexer())
}
