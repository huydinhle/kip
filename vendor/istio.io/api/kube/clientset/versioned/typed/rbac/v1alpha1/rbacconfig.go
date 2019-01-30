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

package v1alpha1

import (
	v1alpha1 "istio.io/api/kube/apis/rbac/v1alpha1"
	scheme "istio.io/api/kube/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RbacConfigsGetter has a method to return a RbacConfigInterface.
// A group's client should implement this interface.
type RbacConfigsGetter interface {
	RbacConfigs(namespace string) RbacConfigInterface
}

// RbacConfigInterface has methods to work with RbacConfig resources.
type RbacConfigInterface interface {
	Create(*v1alpha1.RbacConfig) (*v1alpha1.RbacConfig, error)
	Update(*v1alpha1.RbacConfig) (*v1alpha1.RbacConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.RbacConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.RbacConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RbacConfig, err error)
	RbacConfigExpansion
}

// rbacConfigs implements RbacConfigInterface
type rbacConfigs struct {
	client rest.Interface
	ns     string
}

// newRbacConfigs returns a RbacConfigs
func newRbacConfigs(c *RbacV1alpha1Client, namespace string) *rbacConfigs {
	return &rbacConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rbacConfig, and returns the corresponding rbacConfig object, and an error if there is any.
func (c *rbacConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.RbacConfig, err error) {
	result = &v1alpha1.RbacConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rbacconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RbacConfigs that match those selectors.
func (c *rbacConfigs) List(opts v1.ListOptions) (result *v1alpha1.RbacConfigList, err error) {
	result = &v1alpha1.RbacConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rbacconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rbacConfigs.
func (c *rbacConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rbacconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a rbacConfig and creates it.  Returns the server's representation of the rbacConfig, and an error, if there is any.
func (c *rbacConfigs) Create(rbacConfig *v1alpha1.RbacConfig) (result *v1alpha1.RbacConfig, err error) {
	result = &v1alpha1.RbacConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rbacconfigs").
		Body(rbacConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a rbacConfig and updates it. Returns the server's representation of the rbacConfig, and an error, if there is any.
func (c *rbacConfigs) Update(rbacConfig *v1alpha1.RbacConfig) (result *v1alpha1.RbacConfig, err error) {
	result = &v1alpha1.RbacConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rbacconfigs").
		Name(rbacConfig.Name).
		Body(rbacConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the rbacConfig and deletes it. Returns an error if one occurs.
func (c *rbacConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rbacconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rbacConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rbacconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched rbacConfig.
func (c *rbacConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.RbacConfig, err error) {
	result = &v1alpha1.RbacConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rbacconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
