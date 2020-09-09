// Code generated by xns-informer-gen. DO NOT EDIT.
package core

import (
	xnsinformers "github.com/maistra/xns-informer/pkg/informers"
	v1 "github.com/maistra/xns-informer/pkg/informers/kube/core/v1"
)

type Interface interface {
	V1() v1.Interface
}

type group struct {
	factory xnsinformers.InformerFactory
}

func New(factory xnsinformers.InformerFactory) Interface {
	return &group{factory: factory}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.factory)
}
