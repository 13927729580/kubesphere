/*
Copyright 2020 The KubeSphere Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	calicov3 "kubesphere.io/kubesphere/pkg/apis/network/calicov3"
)

// FakeIPAMBlocks implements IPAMBlockInterface
type FakeIPAMBlocks struct {
	Fake *FakeCrdCalicov3
}

var ipamblocksResource = schema.GroupVersionResource{Group: "crd.projectcalico.org", Version: "calicov3", Resource: "ipamblocks"}

var ipamblocksKind = schema.GroupVersionKind{Group: "crd.projectcalico.org", Version: "calicov3", Kind: "IPAMBlock"}

// Get takes name of the iPAMBlock, and returns the corresponding iPAMBlock object, and an error if there is any.
func (c *FakeIPAMBlocks) Get(ctx context.Context, name string, options v1.GetOptions) (result *calicov3.IPAMBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ipamblocksResource, name), &calicov3.IPAMBlock{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPAMBlock), err
}

// List takes label and field selectors, and returns the list of IPAMBlocks that match those selectors.
func (c *FakeIPAMBlocks) List(ctx context.Context, opts v1.ListOptions) (result *calicov3.IPAMBlockList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ipamblocksResource, ipamblocksKind, opts), &calicov3.IPAMBlockList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &calicov3.IPAMBlockList{ListMeta: obj.(*calicov3.IPAMBlockList).ListMeta}
	for _, item := range obj.(*calicov3.IPAMBlockList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested iPAMBlocks.
func (c *FakeIPAMBlocks) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ipamblocksResource, opts))
}

// Create takes the representation of a iPAMBlock and creates it.  Returns the server's representation of the iPAMBlock, and an error, if there is any.
func (c *FakeIPAMBlocks) Create(ctx context.Context, iPAMBlock *calicov3.IPAMBlock, opts v1.CreateOptions) (result *calicov3.IPAMBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ipamblocksResource, iPAMBlock), &calicov3.IPAMBlock{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPAMBlock), err
}

// Update takes the representation of a iPAMBlock and updates it. Returns the server's representation of the iPAMBlock, and an error, if there is any.
func (c *FakeIPAMBlocks) Update(ctx context.Context, iPAMBlock *calicov3.IPAMBlock, opts v1.UpdateOptions) (result *calicov3.IPAMBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ipamblocksResource, iPAMBlock), &calicov3.IPAMBlock{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPAMBlock), err
}

// Delete takes name of the iPAMBlock and deletes it. Returns an error if one occurs.
func (c *FakeIPAMBlocks) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(ipamblocksResource, name), &calicov3.IPAMBlock{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIPAMBlocks) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ipamblocksResource, listOpts)

	_, err := c.Fake.Invokes(action, &calicov3.IPAMBlockList{})
	return err
}

// Patch applies the patch and returns the patched iPAMBlock.
func (c *FakeIPAMBlocks) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *calicov3.IPAMBlock, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ipamblocksResource, name, pt, data, subresources...), &calicov3.IPAMBlock{})
	if obj == nil {
		return nil, err
	}
	return obj.(*calicov3.IPAMBlock), err
}
