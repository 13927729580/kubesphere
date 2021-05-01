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

// FederatedStatefulSetsGetter has a method to return a FederatedStatefulSetInterface.
// A group's client should implement this interface.
type FederatedStatefulSetsGetter interface {
	FederatedStatefulSets(namespace string) FederatedStatefulSetInterface
}

// FederatedStatefulSetInterface has methods to work with FederatedStatefulSet resources.
type FederatedStatefulSetInterface interface {
	Create(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.CreateOptions) (*v1beta1.FederatedStatefulSet, error)
	Update(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.UpdateOptions) (*v1beta1.FederatedStatefulSet, error)
	UpdateStatus(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.UpdateOptions) (*v1beta1.FederatedStatefulSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.FederatedStatefulSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.FederatedStatefulSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedStatefulSet, err error)
	FederatedStatefulSetExpansion
}

// federatedStatefulSets implements FederatedStatefulSetInterface
type federatedStatefulSets struct {
	client rest.Interface
	ns     string
}

// newFederatedStatefulSets returns a FederatedStatefulSets
func newFederatedStatefulSets(c *TypesV1beta1Client, namespace string) *federatedStatefulSets {
	return &federatedStatefulSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the federatedStatefulSet, and returns the corresponding federatedStatefulSet object, and an error if there is any.
func (c *federatedStatefulSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedStatefulSet, err error) {
	result = &v1beta1.FederatedStatefulSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of FederatedStatefulSets that match those selectors.
func (c *federatedStatefulSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedStatefulSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.FederatedStatefulSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested federatedStatefulSets.
func (c *federatedStatefulSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a federatedStatefulSet and creates it.  Returns the server's representation of the federatedStatefulSet, and an error, if there is any.
func (c *federatedStatefulSets) Create(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.CreateOptions) (result *v1beta1.FederatedStatefulSet, err error) {
	result = &v1beta1.FederatedStatefulSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedStatefulSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a federatedStatefulSet and updates it. Returns the server's representation of the federatedStatefulSet, and an error, if there is any.
func (c *federatedStatefulSets) Update(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.UpdateOptions) (result *v1beta1.FederatedStatefulSet, err error) {
	result = &v1beta1.FederatedStatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		Name(federatedStatefulSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedStatefulSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *federatedStatefulSets) UpdateStatus(ctx context.Context, federatedStatefulSet *v1beta1.FederatedStatefulSet, opts v1.UpdateOptions) (result *v1beta1.FederatedStatefulSet, err error) {
	result = &v1beta1.FederatedStatefulSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		Name(federatedStatefulSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(federatedStatefulSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the federatedStatefulSet and deletes it. Returns an error if one occurs.
func (c *federatedStatefulSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *federatedStatefulSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched federatedStatefulSet.
func (c *federatedStatefulSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedStatefulSet, err error) {
	result = &v1beta1.FederatedStatefulSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("federatedstatefulsets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
