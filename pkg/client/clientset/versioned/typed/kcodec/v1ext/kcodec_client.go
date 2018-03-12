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
package v1ext

import (
	v1ext "github.com/dbenque/kcodec/pkg/api/kcodec/v1ext"
	"github.com/dbenque/kcodec/pkg/client/clientset/versioned/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type KcodecV1extInterface interface {
	RESTClient() rest.Interface
	ItemsGetter
}

// KcodecV1extClient is used to interact with features provided by the kcodec group.
type KcodecV1extClient struct {
	restClient rest.Interface
}

func (c *KcodecV1extClient) Items(namespace string) ItemInterface {
	return newItems(c, namespace)
}

// NewForConfig creates a new KcodecV1extClient for the given config.
func NewForConfig(c *rest.Config) (*KcodecV1extClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &KcodecV1extClient{client}, nil
}

// NewForConfigOrDie creates a new KcodecV1extClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *KcodecV1extClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new KcodecV1extClient for the given RESTClient.
func New(c rest.Interface) *KcodecV1extClient {
	return &KcodecV1extClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1ext.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *KcodecV1extClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
