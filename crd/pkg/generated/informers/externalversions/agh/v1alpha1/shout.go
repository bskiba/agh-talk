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

	aghv1alpha1 "github.com/bskiba/agh-talk/crd/pkg/apis/agh/v1alpha1"
	versioned "github.com/bskiba/agh-talk/crd/pkg/generated/clientset/versioned"
	internalinterfaces "github.com/bskiba/agh-talk/crd/pkg/generated/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/bskiba/agh-talk/crd/pkg/generated/listers/agh/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ShoutInformer provides access to a shared informer and lister for
// Shouts.
type ShoutInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ShoutLister
}

type shoutInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewShoutInformer constructs a new informer for Shout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewShoutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredShoutInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredShoutInformer constructs a new informer for Shout type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredShoutInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AghV1alpha1().Shouts(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.AghV1alpha1().Shouts(namespace).Watch(options)
			},
		},
		&aghv1alpha1.Shout{},
		resyncPeriod,
		indexers,
	)
}

func (f *shoutInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredShoutInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *shoutInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&aghv1alpha1.Shout{}, f.defaultInformer)
}

func (f *shoutInformer) Lister() v1alpha1.ShoutLister {
	return v1alpha1.NewShoutLister(f.Informer().GetIndexer())
}