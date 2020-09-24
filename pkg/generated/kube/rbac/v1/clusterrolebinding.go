// Code generated by xns-informer-gen. DO NOT EDIT.

package v1

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "k8s.io/api/rbac/v1"
	informers "k8s.io/client-go/informers/rbac/v1"
	listers "k8s.io/client-go/listers/rbac/v1"
	"k8s.io/client-go/tools/cache"
)

type clusterRoleBindingInformer struct {
	informer cache.SharedIndexInformer
}

var _ informers.ClusterRoleBindingInformer = &clusterRoleBindingInformer{}

func NewClusterRoleBindingInformer(f xnsinformers.SharedInformerFactory) informers.ClusterRoleBindingInformer {
	resource := v1.SchemeGroupVersion.WithResource("clusterrolebindings")
	informer := f.ClusterResource(resource).Informer()

	return &clusterRoleBindingInformer{
		informer: xnsinformers.NewInformerConverter(f.GetScheme(), informer, &v1.ClusterRoleBinding{}),
	}
}

func (i *clusterRoleBindingInformer) Informer() cache.SharedIndexInformer {
	return i.informer
}

func (i *clusterRoleBindingInformer) Lister() listers.ClusterRoleBindingLister {
	return listers.NewClusterRoleBindingLister(i.informer.GetIndexer())
}