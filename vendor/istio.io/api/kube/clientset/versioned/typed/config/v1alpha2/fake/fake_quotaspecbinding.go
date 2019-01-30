// Copyright 2019 Istio Authors
//
//   Licensed under the Apache License, Version 2.0 (the "License");
//   you may not use this file except in compliance with the License.
//   You may obtain a copy of the License at
//
//       http://www.apache.org/licenses/LICENSE-2.0
//
//   Unless required by applicable law or agreed to in writing, software
//   distributed under the License is distributed on an "AS IS" BASIS,
//   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//   See the License for the specific language governing permissions and
//   limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha2 "istio.io/api/kube/apis/config/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeQuotaSpecBindings implements QuotaSpecBindingInterface
type FakeQuotaSpecBindings struct {
	Fake *FakeConfigV1alpha2
	ns   string
}

var quotaspecbindingsResource = schema.GroupVersionResource{Group: "config", Version: "v1alpha2", Resource: "quotaspecbindings"}

var quotaspecbindingsKind = schema.GroupVersionKind{Group: "config", Version: "v1alpha2", Kind: "QuotaSpecBinding"}

// Get takes name of the quotaSpecBinding, and returns the corresponding quotaSpecBinding object, and an error if there is any.
func (c *FakeQuotaSpecBindings) Get(name string, options v1.GetOptions) (result *v1alpha2.QuotaSpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(quotaspecbindingsResource, c.ns, name), &v1alpha2.QuotaSpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.QuotaSpecBinding), err
}

// List takes label and field selectors, and returns the list of QuotaSpecBindings that match those selectors.
func (c *FakeQuotaSpecBindings) List(opts v1.ListOptions) (result *v1alpha2.QuotaSpecBindingList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(quotaspecbindingsResource, quotaspecbindingsKind, c.ns, opts), &v1alpha2.QuotaSpecBindingList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.QuotaSpecBindingList{ListMeta: obj.(*v1alpha2.QuotaSpecBindingList).ListMeta}
	for _, item := range obj.(*v1alpha2.QuotaSpecBindingList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested quotaSpecBindings.
func (c *FakeQuotaSpecBindings) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(quotaspecbindingsResource, c.ns, opts))

}

// Create takes the representation of a quotaSpecBinding and creates it.  Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *FakeQuotaSpecBindings) Create(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(quotaspecbindingsResource, c.ns, quotaSpecBinding), &v1alpha2.QuotaSpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.QuotaSpecBinding), err
}

// Update takes the representation of a quotaSpecBinding and updates it. Returns the server's representation of the quotaSpecBinding, and an error, if there is any.
func (c *FakeQuotaSpecBindings) Update(quotaSpecBinding *v1alpha2.QuotaSpecBinding) (result *v1alpha2.QuotaSpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(quotaspecbindingsResource, c.ns, quotaSpecBinding), &v1alpha2.QuotaSpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.QuotaSpecBinding), err
}

// Delete takes name of the quotaSpecBinding and deletes it. Returns an error if one occurs.
func (c *FakeQuotaSpecBindings) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(quotaspecbindingsResource, c.ns, name), &v1alpha2.QuotaSpecBinding{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeQuotaSpecBindings) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(quotaspecbindingsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.QuotaSpecBindingList{})
	return err
}

// Patch applies the patch and returns the patched quotaSpecBinding.
func (c *FakeQuotaSpecBindings) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.QuotaSpecBinding, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(quotaspecbindingsResource, c.ns, name, data, subresources...), &v1alpha2.QuotaSpecBinding{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.QuotaSpecBinding), err
}
