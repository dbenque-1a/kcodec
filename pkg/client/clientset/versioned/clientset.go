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
package versioned

import (
	kcodecv1 "github.com/dbenque/kcodec/pkg/client/clientset/versioned/typed/kcodec/v1"
	kcodecv1ext "github.com/dbenque/kcodec/pkg/client/clientset/versioned/typed/kcodec/v1ext"
	kcodecv2 "github.com/dbenque/kcodec/pkg/client/clientset/versioned/typed/kcodec/v2"
	glog "github.com/golang/glog"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	KcodecV1() kcodecv1.KcodecV1Interface
	KcodecV1ext() kcodecv1ext.KcodecV1extInterface
	KcodecV2() kcodecv2.KcodecV2Interface
	// Deprecated: please explicitly pick a version if possible.
	Kcodec() kcodecv2.KcodecV2Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	kcodecV1    *kcodecv1.KcodecV1Client
	kcodecV1ext *kcodecv1ext.KcodecV1extClient
	kcodecV2    *kcodecv2.KcodecV2Client
}

// KcodecV1 retrieves the KcodecV1Client
func (c *Clientset) KcodecV1() kcodecv1.KcodecV1Interface {
	return c.kcodecV1
}

// KcodecV1ext retrieves the KcodecV1extClient
func (c *Clientset) KcodecV1ext() kcodecv1ext.KcodecV1extInterface {
	return c.kcodecV1ext
}

// KcodecV2 retrieves the KcodecV2Client
func (c *Clientset) KcodecV2() kcodecv2.KcodecV2Interface {
	return c.kcodecV2
}

// Deprecated: Kcodec retrieves the default version of KcodecClient.
// Please explicitly pick a version.
func (c *Clientset) Kcodec() kcodecv2.KcodecV2Interface {
	return c.kcodecV2
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.kcodecV1, err = kcodecv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kcodecV1ext, err = kcodecv1ext.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.kcodecV2, err = kcodecv2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		glog.Errorf("failed to create the DiscoveryClient: %v", err)
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.kcodecV1 = kcodecv1.NewForConfigOrDie(c)
	cs.kcodecV1ext = kcodecv1ext.NewForConfigOrDie(c)
	cs.kcodecV2 = kcodecv2.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.kcodecV1 = kcodecv1.New(c)
	cs.kcodecV1ext = kcodecv1ext.New(c)
	cs.kcodecV2 = kcodecv2.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
