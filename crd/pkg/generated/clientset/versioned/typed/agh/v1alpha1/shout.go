/*
Copyright The Kubernetes Authors.

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

package v1alpha1

import (
	"time"

	v1alpha1 "github.com/bskiba/agh-talk/crd/pkg/apis/agh/v1alpha1"
	scheme "github.com/bskiba/agh-talk/crd/pkg/generated/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ShoutsGetter has a method to return a ShoutInterface.
// A group's client should implement this interface.
type ShoutsGetter interface {
	Shouts(namespace string) ShoutInterface
}

// ShoutInterface has methods to work with Shout resources.
type ShoutInterface interface {
	Create(*v1alpha1.Shout) (*v1alpha1.Shout, error)
	Update(*v1alpha1.Shout) (*v1alpha1.Shout, error)
	UpdateStatus(*v1alpha1.Shout) (*v1alpha1.Shout, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Shout, error)
	List(opts v1.ListOptions) (*v1alpha1.ShoutList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Shout, err error)
	ShoutExpansion
}

// shouts implements ShoutInterface
type shouts struct {
	client rest.Interface
	ns     string
}

// newShouts returns a Shouts
func newShouts(c *AghV1alpha1Client, namespace string) *shouts {
	return &shouts{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the shout, and returns the corresponding shout object, and an error if there is any.
func (c *shouts) Get(name string, options v1.GetOptions) (result *v1alpha1.Shout, err error) {
	result = &v1alpha1.Shout{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shouts").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Shouts that match those selectors.
func (c *shouts) List(opts v1.ListOptions) (result *v1alpha1.ShoutList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ShoutList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("shouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested shouts.
func (c *shouts) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("shouts").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a shout and creates it.  Returns the server's representation of the shout, and an error, if there is any.
func (c *shouts) Create(shout *v1alpha1.Shout) (result *v1alpha1.Shout, err error) {
	result = &v1alpha1.Shout{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("shouts").
		Body(shout).
		Do().
		Into(result)
	return
}

// Update takes the representation of a shout and updates it. Returns the server's representation of the shout, and an error, if there is any.
func (c *shouts) Update(shout *v1alpha1.Shout) (result *v1alpha1.Shout, err error) {
	result = &v1alpha1.Shout{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shouts").
		Name(shout.Name).
		Body(shout).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *shouts) UpdateStatus(shout *v1alpha1.Shout) (result *v1alpha1.Shout, err error) {
	result = &v1alpha1.Shout{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("shouts").
		Name(shout.Name).
		SubResource("status").
		Body(shout).
		Do().
		Into(result)
	return
}

// Delete takes name of the shout and deletes it. Returns an error if one occurs.
func (c *shouts) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shouts").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *shouts) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("shouts").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched shout.
func (c *shouts) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Shout, err error) {
	result = &v1alpha1.Shout{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("shouts").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
