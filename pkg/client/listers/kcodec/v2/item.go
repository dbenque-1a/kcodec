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

// This file was automatically generated by lister-gen

package v2

import (
	v2 "github.com/dbenque/kcodec/pkg/api/kcodec/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ItemLister helps list Items.
type ItemLister interface {
	// List lists all Items in the indexer.
	List(selector labels.Selector) (ret []*v2.Item, err error)
	// Items returns an object that can list and get Items.
	Items(namespace string) ItemNamespaceLister
	ItemListerExpansion
}

// itemLister implements the ItemLister interface.
type itemLister struct {
	indexer cache.Indexer
}

// NewItemLister returns a new ItemLister.
func NewItemLister(indexer cache.Indexer) ItemLister {
	return &itemLister{indexer: indexer}
}

// List lists all Items in the indexer.
func (s *itemLister) List(selector labels.Selector) (ret []*v2.Item, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Item))
	})
	return ret, err
}

// Items returns an object that can list and get Items.
func (s *itemLister) Items(namespace string) ItemNamespaceLister {
	return itemNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ItemNamespaceLister helps list and get Items.
type ItemNamespaceLister interface {
	// List lists all Items in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v2.Item, err error)
	// Get retrieves the Item from the indexer for a given namespace and name.
	Get(name string) (*v2.Item, error)
	ItemNamespaceListerExpansion
}

// itemNamespaceLister implements the ItemNamespaceLister
// interface.
type itemNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Items in the indexer for a given namespace.
func (s itemNamespaceLister) List(selector labels.Selector) (ret []*v2.Item, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.Item))
	})
	return ret, err
}

// Get retrieves the Item from the indexer for a given namespace and name.
func (s itemNamespaceLister) Get(name string) (*v2.Item, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("item"), name)
	}
	return obj.(*v2.Item), nil
}
