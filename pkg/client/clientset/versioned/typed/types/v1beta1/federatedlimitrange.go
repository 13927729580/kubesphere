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

package v1beta1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// FederatedLimitRangesGetter has a method to return a FederatedLimitRangeInterface.
// A group's client should implement this interface.
type FederatedLimitRangesGetter interface {
	FederatedLimitRanges(namespace string) FederatedLimitRangeInterface
}

// FederatedLimitRangeInterface has methods to work with FederatedLimitRange resources.
type FederatedLimitRangeInterface interface {
	Create(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.CreateOptions) (*v1beta1.FederatedLimitRange, error)
	Update(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.UpdateOptions) (*v1beta1.FederatedLimitRange, error)
	UpdateStatus(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.UpdateOptions) (*v1beta1.FederatedLimitRange, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FederatedLimitRange, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FederatedLimitRangeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedLimitRange, err error)
	FederatedLimitRangeExpansion
}

// federatedLimitRanges implements FederatedLimitRangeInterface
type federatedLimitRanges struct {
	client rest.Interface
	ns     string
}

// newFederatedLimitRanges returns a FederatedLimitRanges
func newFederatedLimitRanges(c *TypesV1beta1Client, namespace string) *federatedLimitRanges {
	return &federatedLimitRanges{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedLimitRange, and returns the corresponding federatedLimitRange object, and an error if there is any.
func (c *federatedLimitRanges) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedLimitRange, err error) {
	result = &v1beta1.FederatedLimitRange{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedLimitRanges that match those selectors.
func (c *federatedLimitRanges) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedLimitRangeList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedLimitRangeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedLimitRanges.
func (c *federatedLimitRanges) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedLimitRange and creates it.  Returns the server's representation of the federatedLimitRange, and an error, if there is any.
func (c *federatedLimitRanges) Create(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.CreateOptions) (result *v1beta1.FederatedLimitRange, err error) {
	result = &v1beta1.FederatedLimitRange{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedLimitRange).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedLimitRange and updates it. Returns the server's representation of the federatedLimitRange, and an error, if there is any.
func (c *federatedLimitRanges) Update(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.UpdateOptions) (result *v1beta1.FederatedLimitRange, err error) {
	result = &v1beta1.FederatedLimitRange{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		Name(federatedLimitRange.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedLimitRange).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedLimitRanges) UpdateStatus(ctx context.Context, federatedLimitRange *v1beta1.FederatedLimitRange, opts v1.UpdateOptions) (result *v1beta1.FederatedLimitRange, err error) {
	result = &v1beta1.FederatedLimitRange{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		Name(federatedLimitRange.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedLimitRange).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedLimitRange and deletes it. Returns an error if one occurs.
func (c *federatedLimitRanges) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedLimitRanges) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedlimitranges").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedLimitRange.
func (c *federatedLimitRanges) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedLimitRange, err error) {
	result = &v1beta1.FederatedLimitRange{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedlimitranges").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
