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

// FederatedWorkspaceInformer provides access to a shared informer and lister for
// FederatedWorkspaces.
type FederatedWorkspaceInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1beta1.FederatedWorkspaceLister
}

type federatedWorkspaceInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewFederatedWorkspaceInformer constructs a new informer for FederatedWorkspace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFederatedWorkspaceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredFederatedWorkspaceInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredFederatedWorkspaceInformer constructs a new informer for FederatedWorkspace type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredFederatedWorkspaceInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedWorkspaces().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TypesV1beta1().FederatedWorkspaces().Watch(context.TODO(), options)
			},
		},
		&typesv1beta1.FederatedWorkspace{},
		resyncPeriod,
		indexers,
	)
}

func (f *federatedWorkspaceInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredFederatedWorkspaceInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *federatedWorkspaceInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&typesv1beta1.FederatedWorkspace{}, f.defaultInformer)
}

func (f *federatedWorkspaceInformer) Lister() v1beta1.FederatedWorkspaceLister {
	return v1beta1.NewFederatedWorkspaceLister(f.Informer().GetIndexer())
}
