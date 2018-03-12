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
	scheme "github.com/dbenque/kcodec/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ItemsGetter has a method to return a ItemInterface.
// A group's client should implement this interface.
type ItemsGetter interface {
	Items(namespace string) ItemInterface
}

// ItemInterface has methods to work with Item resources.
type ItemInterface interface {
	Create(*v1ext.Item) (*v1ext.Item, error)
	Update(*v1ext.Item) (*v1ext.Item, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1ext.Item, error)
	List(opts v1.ListOptions) (*v1ext.ItemList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1ext.Item, err error)
	ItemExpansion
}

// items implements ItemInterface
type items struct {
	client rest.Interface
	ns     string
}

// newItems returns a Items
func newItems(c *KcodecV1extClient, namespace string) *items {
	return &items{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the item, and returns the corresponding item object, and an error if there is any.
func (c *items) Get(name string, options v1.GetOptions) (result *v1ext.Item, err error) {
	result = &v1ext.Item{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("items").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Items that match those selectors.
func (c *items) List(opts v1.ListOptions) (result *v1ext.ItemList, err error) {
	result = &v1ext.ItemList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("items").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested items.
func (c *items) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("items").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a item and creates it.  Returns the server's representation of the item, and an error, if there is any.
func (c *items) Create(item *v1ext.Item) (result *v1ext.Item, err error) {
	result = &v1ext.Item{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("items").
		Body(item).
		Do().
		Into(result)
	return
}

// Update takes the representation of a item and updates it. Returns the server's representation of the item, and an error, if there is any.
func (c *items) Update(item *v1ext.Item) (result *v1ext.Item, err error) {
	result = &v1ext.Item{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("items").
		Name(item.Name).
		Body(item).
		Do().
		Into(result)
	return
}

// Delete takes name of the item and deletes it. Returns an error if one occurs.
func (c *items) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("items").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *items) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("items").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched item.
func (c *items) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1ext.Item, err error) {
	result = &v1ext.Item{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("items").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
