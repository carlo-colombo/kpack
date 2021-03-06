/*
 * Copyright 2019 The original author or authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/pivotal/kpack/pkg/apis/experimental/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStores implements StoreInterface
type FakeStores struct {
	Fake *FakeExperimentalV1alpha1
}

var storesResource = schema.GroupVersionResource{Group: "experimental.kpack.pivotal.io", Version: "v1alpha1", Resource: "stores"}

var storesKind = schema.GroupVersionKind{Group: "experimental.kpack.pivotal.io", Version: "v1alpha1", Kind: "Store"}

// Get takes name of the store, and returns the corresponding store object, and an error if there is any.
func (c *FakeStores) Get(name string, options v1.GetOptions) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(storesResource, name), &v1alpha1.Store{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// List takes label and field selectors, and returns the list of Stores that match those selectors.
func (c *FakeStores) List(opts v1.ListOptions) (result *v1alpha1.StoreList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(storesResource, storesKind, opts), &v1alpha1.StoreList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.StoreList{ListMeta: obj.(*v1alpha1.StoreList).ListMeta}
	for _, item := range obj.(*v1alpha1.StoreList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested stores.
func (c *FakeStores) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(storesResource, opts))
}

// Create takes the representation of a store and creates it.  Returns the server's representation of the store, and an error, if there is any.
func (c *FakeStores) Create(store *v1alpha1.Store) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(storesResource, store), &v1alpha1.Store{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// Update takes the representation of a store and updates it. Returns the server's representation of the store, and an error, if there is any.
func (c *FakeStores) Update(store *v1alpha1.Store) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(storesResource, store), &v1alpha1.Store{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStores) UpdateStatus(store *v1alpha1.Store) (*v1alpha1.Store, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(storesResource, "status", store), &v1alpha1.Store{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}

// Delete takes name of the store and deletes it. Returns an error if one occurs.
func (c *FakeStores) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(storesResource, name), &v1alpha1.Store{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStores) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(storesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.StoreList{})
	return err
}

// Patch applies the patch and returns the patched store.
func (c *FakeStores) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Store, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(storesResource, name, pt, data, subresources...), &v1alpha1.Store{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Store), err
}
