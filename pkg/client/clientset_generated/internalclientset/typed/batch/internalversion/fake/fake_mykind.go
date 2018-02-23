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
package fake

import (
	batch "github.com/the-redback/apiserver-builded/pkg/apis/batch"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMyKinds implements MyKindInterface
type FakeMyKinds struct {
	Fake *FakeBatch
	ns   string
}

var mykindsResource = schema.GroupVersionResource{Group: "batch.iammaruf.com", Version: "", Resource: "mykinds"}

var mykindsKind = schema.GroupVersionKind{Group: "batch.iammaruf.com", Version: "", Kind: "MyKind"}

// Get takes name of the myKind, and returns the corresponding myKind object, and an error if there is any.
func (c *FakeMyKinds) Get(name string, options v1.GetOptions) (result *batch.MyKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(mykindsResource, c.ns, name), &batch.MyKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.MyKind), err
}

// List takes label and field selectors, and returns the list of MyKinds that match those selectors.
func (c *FakeMyKinds) List(opts v1.ListOptions) (result *batch.MyKindList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(mykindsResource, mykindsKind, c.ns, opts), &batch.MyKindList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &batch.MyKindList{}
	for _, item := range obj.(*batch.MyKindList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested myKinds.
func (c *FakeMyKinds) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(mykindsResource, c.ns, opts))

}

// Create takes the representation of a myKind and creates it.  Returns the server's representation of the myKind, and an error, if there is any.
func (c *FakeMyKinds) Create(myKind *batch.MyKind) (result *batch.MyKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(mykindsResource, c.ns, myKind), &batch.MyKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.MyKind), err
}

// Update takes the representation of a myKind and updates it. Returns the server's representation of the myKind, and an error, if there is any.
func (c *FakeMyKinds) Update(myKind *batch.MyKind) (result *batch.MyKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(mykindsResource, c.ns, myKind), &batch.MyKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.MyKind), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMyKinds) UpdateStatus(myKind *batch.MyKind) (*batch.MyKind, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(mykindsResource, "status", c.ns, myKind), &batch.MyKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.MyKind), err
}

// Delete takes name of the myKind and deletes it. Returns an error if one occurs.
func (c *FakeMyKinds) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(mykindsResource, c.ns, name), &batch.MyKind{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMyKinds) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(mykindsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &batch.MyKindList{})
	return err
}

// Patch applies the patch and returns the patched myKind.
func (c *FakeMyKinds) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *batch.MyKind, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(mykindsResource, c.ns, name, data, subresources...), &batch.MyKind{})

	if obj == nil {
		return nil, err
	}
	return obj.(*batch.MyKind), err
}
