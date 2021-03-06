/*
Copyright 2017 The Kubernetes Authors.

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
package internalversion

import (
	batch "github.com/the-redback/apiserver-builded/pkg/apis/batch"
	scheme "github.com/the-redback/apiserver-builded/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// MyKindsGetter has a method to return a MyKindInterface.
// A group's client should implement this interface.
type MyKindsGetter interface {
	MyKinds(namespace string) MyKindInterface
}

// MyKindInterface has methods to work with MyKind resources.
type MyKindInterface interface {
	Create(*batch.MyKind) (*batch.MyKind, error)
	Update(*batch.MyKind) (*batch.MyKind, error)
	UpdateStatus(*batch.MyKind) (*batch.MyKind, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*batch.MyKind, error)
	List(opts v1.ListOptions) (*batch.MyKindList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.MyKind, err error)
	MyKindExpansion
}

// myKinds implements MyKindInterface
type myKinds struct {
	client rest.Interface
	ns     string
}

// newMyKinds returns a MyKinds
func newMyKinds(c *BatchClient, namespace string) *myKinds {
	return &myKinds{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the myKind, and returns the corresponding myKind object, and an error if there is any.
func (c *myKinds) Get(name string, options v1.GetOptions) (result *batch.MyKind, err error) {
	result = &batch.MyKind{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of MyKinds that match those selectors.
func (c *myKinds) List(opts v1.ListOptions) (result *batch.MyKindList, err error) {
	result = &batch.MyKindList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested myKinds.
func (c *myKinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a myKind and creates it.  Returns the server's representation of the myKind, and an error, if there is any.
func (c *myKinds) Create(myKind *batch.MyKind) (result *batch.MyKind, err error) {
	result = &batch.MyKind{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("mykinds").
		Body(myKind).
		Do().
		Into(result)
	return
}

// Update takes the representation of a myKind and updates it. Returns the server's representation of the myKind, and an error, if there is any.
func (c *myKinds) Update(myKind *batch.MyKind) (result *batch.MyKind, err error) {
	result = &batch.MyKind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mykinds").
		Name(myKind.Name).
		Body(myKind).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *myKinds) UpdateStatus(myKind *batch.MyKind) (result *batch.MyKind, err error) {
	result = &batch.MyKind{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("mykinds").
		Name(myKind.Name).
		SubResource("status").
		Body(myKind).
		Do().
		Into(result)
	return
}

// Delete takes name of the myKind and deletes it. Returns an error if one occurs.
func (c *myKinds) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mykinds").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *myKinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("mykinds").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched myKind.
func (c *myKinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.MyKind, err error) {
	result = &batch.MyKind{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("mykinds").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
