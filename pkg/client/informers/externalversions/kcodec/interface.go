/*
MIT License

Copyright (c) 2018 PodKubervisor

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// This file was automatically generated by informer-gen

package kcodec

import (
	internalinterfaces "github.com/dbenque/kcodec/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/dbenque/kcodec/pkg/client/informers/externalversions/kcodec/v1"
	v1ext "github.com/dbenque/kcodec/pkg/client/informers/externalversions/kcodec/v1ext"
	v2 "github.com/dbenque/kcodec/pkg/client/informers/externalversions/kcodec/v2"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V1 provides access to shared informers for resources in V1.
	V1() v1.Interface
	// V1ext provides access to shared informers for resources in V1ext.
	V1ext() v1ext.Interface
	// V2 provides access to shared informers for resources in V2.
	V2() v2.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V1 returns a new v1.Interface.
func (g *group) V1() v1.Interface {
	return v1.New(g.factory, g.namespace, g.tweakListOptions)
}

// V1ext returns a new v1ext.Interface.
func (g *group) V1ext() v1ext.Interface {
	return v1ext.New(g.factory, g.namespace, g.tweakListOptions)
}

// V2 returns a new v2.Interface.
func (g *group) V2() v2.Interface {
	return v2.New(g.factory, g.namespace, g.tweakListOptions)
}
