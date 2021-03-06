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
package fake

import (
	kcodec_v1 "github.com/dbenque/kcodec/pkg/api/kcodec/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeItems implements ItemInterface
type FakeItems struct {
	Fake *FakeKcodecV1
	ns   string
}

var itemsResource = schema.GroupVersionResource{Group: "kcodec", Version: "v1", Resource: "items"}

var itemsKind = schema.GroupVersionKind{Group: "kcodec", Version: "v1", Kind: "Item"}

// Get takes name of the item, and returns the corresponding item object, and an error if there is any.
func (c *FakeItems) Get(name string, options v1.GetOptions) (result *kcodec_v1.Item, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(itemsResource, c.ns, name), &kcodec_v1.Item{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kcodec_v1.Item), err
}

// List takes label and field selectors, and returns the list of Items that match those selectors.
func (c *FakeItems) List(opts v1.ListOptions) (result *kcodec_v1.ItemList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(itemsResource, itemsKind, c.ns, opts), &kcodec_v1.ItemList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &kcodec_v1.ItemList{}
	for _, item := range obj.(*kcodec_v1.ItemList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested items.
func (c *FakeItems) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(itemsResource, c.ns, opts))

}

// Create takes the representation of a item and creates it.  Returns the server's representation of the item, and an error, if there is any.
func (c *FakeItems) Create(item *kcodec_v1.Item) (result *kcodec_v1.Item, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(itemsResource, c.ns, item), &kcodec_v1.Item{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kcodec_v1.Item), err
}

// Update takes the representation of a item and updates it. Returns the server's representation of the item, and an error, if there is any.
func (c *FakeItems) Update(item *kcodec_v1.Item) (result *kcodec_v1.Item, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(itemsResource, c.ns, item), &kcodec_v1.Item{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kcodec_v1.Item), err
}

// Delete takes name of the item and deletes it. Returns an error if one occurs.
func (c *FakeItems) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(itemsResource, c.ns, name), &kcodec_v1.Item{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeItems) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(itemsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &kcodec_v1.ItemList{})
	return err
}

// Patch applies the patch and returns the patched item.
func (c *FakeItems) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *kcodec_v1.Item, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(itemsResource, c.ns, name, data, subresources...), &kcodec_v1.Item{})

	if obj == nil {
		return nil, err
	}
	return obj.(*kcodec_v1.Item), err
}
