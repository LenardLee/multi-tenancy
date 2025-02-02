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

	v1alpha1 "github.com/multi-tenancy/incubator/virtualcluster/pkg/apis/tenancy/v1alpha1"
	scheme "github.com/multi-tenancy/incubator/virtualcluster/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// VirtualclustersGetter has a method to return a VirtualclusterInterface.
// A group's client should implement this interface.
type VirtualclustersGetter interface {
	Virtualclusters(namespace string) VirtualclusterInterface
}

// VirtualclusterInterface has methods to work with Virtualcluster resources.
type VirtualclusterInterface interface {
	Create(*v1alpha1.Virtualcluster) (*v1alpha1.Virtualcluster, error)
	Update(*v1alpha1.Virtualcluster) (*v1alpha1.Virtualcluster, error)
	UpdateStatus(*v1alpha1.Virtualcluster) (*v1alpha1.Virtualcluster, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Virtualcluster, error)
	List(opts v1.ListOptions) (*v1alpha1.VirtualclusterList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualcluster, err error)
	VirtualclusterExpansion
}

// virtualclusters implements VirtualclusterInterface
type virtualclusters struct {
	client rest.Interface
	ns     string
}

// newVirtualclusters returns a Virtualclusters
func newVirtualclusters(c *TenancyV1alpha1Client, namespace string) *virtualclusters {
	return &virtualclusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the virtualcluster, and returns the corresponding virtualcluster object, and an error if there is any.
func (c *virtualclusters) Get(name string, options v1.GetOptions) (result *v1alpha1.Virtualcluster, err error) {
	result = &v1alpha1.Virtualcluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Virtualclusters that match those selectors.
func (c *virtualclusters) List(opts v1.ListOptions) (result *v1alpha1.VirtualclusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.VirtualclusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested virtualclusters.
func (c *virtualclusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a virtualcluster and creates it.  Returns the server's representation of the virtualcluster, and an error, if there is any.
func (c *virtualclusters) Create(virtualcluster *v1alpha1.Virtualcluster) (result *v1alpha1.Virtualcluster, err error) {
	result = &v1alpha1.Virtualcluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("virtualclusters").
		Body(virtualcluster).
		Do().
		Into(result)
	return
}

// Update takes the representation of a virtualcluster and updates it. Returns the server's representation of the virtualcluster, and an error, if there is any.
func (c *virtualclusters) Update(virtualcluster *v1alpha1.Virtualcluster) (result *v1alpha1.Virtualcluster, err error) {
	result = &v1alpha1.Virtualcluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(virtualcluster.Name).
		Body(virtualcluster).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *virtualclusters) UpdateStatus(virtualcluster *v1alpha1.Virtualcluster) (result *v1alpha1.Virtualcluster, err error) {
	result = &v1alpha1.Virtualcluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(virtualcluster.Name).
		SubResource("status").
		Body(virtualcluster).
		Do().
		Into(result)
	return
}

// Delete takes name of the virtualcluster and deletes it. Returns an error if one occurs.
func (c *virtualclusters) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualclusters").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *virtualclusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("virtualclusters").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched virtualcluster.
func (c *virtualclusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Virtualcluster, err error) {
	result = &v1alpha1.Virtualcluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("virtualclusters").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
