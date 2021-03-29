/*
Copyright 2020 The KubeSphere Authors.

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

package v1beta1

import (
	"context"
	time "time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"

	typesv1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
	versioned "kubesphere.io/kubesphere/pkg/client/clientset/versioned"
	internalinterfaces "kubesphere.io/kubesphere/pkg/client/informers/externalversions/internalinterfaces"
	v1beta1 "kubesphere.io/kubesphere/pkg/client/listers/types/v1beta1"
)

// FederatedClusterRoleInformer provides access to a shared informer and lister for
// FederatedClusterRoles.
type FederatedClusterRoleInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FederatedClusterRoleLister
}

type federatedClusterRoleInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewFederatedClusterRoleInformer constructs a new informer for FederatedClusterRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedClusterRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedClusterRoleInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedClusterRoleInformer constructs a new informer for FederatedClusterRole type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedClusterRoleInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedClusterRoles(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedClusterRoles(namespace).Watch(context.TODO(), options)
			},
		},
		&typesv1beta1.FederatedClusterRole{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedClusterRoleInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedClusterRoleInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedClusterRoleInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&typesv1beta1.FederatedClusterRole{}, f.defaultInformer)
}

func (f *federatedClusterRoleInformer) Lister() v1beta1.FederatedClusterRoleLister {
	return v1beta1.NewFederatedClusterRoleLister(f.Informer().GetIndexer())
}
