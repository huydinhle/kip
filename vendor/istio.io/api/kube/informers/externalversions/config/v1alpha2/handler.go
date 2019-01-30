// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha2

import (
	time "time"

	configv1alpha2 "istio.io/api/kube/apis/config/v1alpha2"
	versioned "istio.io/api/kube/clientset/versioned"
	internalinterfaces "istio.io/api/kube/informers/externalversions/internalinterfaces"
	v1alpha2 "istio.io/api/kube/listers/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// HandlerInformer provides access to a shared informer and lister for
// Handlers.
type HandlerInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha2.HandlerLister
}

type handlerInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewHandlerInformer constructs a new informer for Handler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewHandlerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredHandlerInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredHandlerInformer constructs a new informer for Handler type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredHandlerInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha2().Handlers(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ConfigV1alpha2().Handlers(namespace).Watch(options)
			},
		},
		&configv1alpha2.Handler{},
		resyncPeriod,
		indexers,
	)
}

func (f *handlerInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredHandlerInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *handlerInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&configv1alpha2.Handler{}, f.defaultInformer)
}

func (f *handlerInformer) Lister() v1alpha2.HandlerLister {
	return v1alpha2.NewHandlerLister(f.Informer().GetIndexer())
}
