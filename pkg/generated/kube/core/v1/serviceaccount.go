// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/core/v1"
	informers "k8s.io/client-go/informers/core/v1"
	listers "k8s.io/client-go/listers/core/v1"
	"k8s.io/client-go/tools/cache"
)

type serviceAccountInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ServiceAccountInformer = &serviceAccountInformer{}

func NewServiceAccountInformer(f xnsinformers.SharedInformerFactory) informers.ServiceAccountInformer {
	resource := v1.SchemeGroupVersion.WithResource("serviceaccounts")
	informer := f.NamespacedResource(resource).Informer()

	return &serviceAccountInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1.ServiceAccount{}),
	}
}

func (i *serviceAccountInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *serviceAccountInformer) Lister() listers.ServiceAccountLister {
	return listers.NewServiceAccountLister(i.informer.GetIndexer())
}