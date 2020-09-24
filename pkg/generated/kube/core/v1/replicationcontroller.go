// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type replicationControllerInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ReplicationControllerInformer = &replicationControllerInformer{}

func NewReplicationControllerInformer(f xnsinformers.SharedInformerFactory) informers.ReplicationControllerInformer {
	resource := v1.SchemeGroupVersion.WithResource("replicationcontrollers")
	informer := f.NamespacedResource(resource).Informer()

	return &replicationControllerInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1.ReplicationController{}),
	}
}

func (i *replicationControllerInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *replicationControllerInformer) Lister() listers.ReplicationControllerLister {
	return listers.NewReplicationControllerLister(i.informer.GetIndexer())
}