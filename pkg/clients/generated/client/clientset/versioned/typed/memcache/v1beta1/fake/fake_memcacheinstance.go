// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/memcache/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeMemcacheInstances implements MemcacheInstanceInterface
type FakeMemcacheInstances struct {
	Fake *FakeMemcacheV1beta1
	ns   string
}

var memcacheinstancesResource = schema.GroupVersionResource{Group: "memcache.cnrm.cloud.google.com", Version: "v1beta1", Resource: "memcacheinstances"}

var memcacheinstancesKind = schema.GroupVersionKind{Group: "memcache.cnrm.cloud.google.com", Version: "v1beta1", Kind: "MemcacheInstance"}

// Get takes name of the memcacheInstance, and returns the corresponding memcacheInstance object, and an error if there is any.
func (c *FakeMemcacheInstances) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.MemcacheInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(memcacheinstancesResource, c.ns, name), &v1beta1.MemcacheInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MemcacheInstance), err
}

// List takes label and field selectors, and returns the list of MemcacheInstances that match those selectors.
func (c *FakeMemcacheInstances) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.MemcacheInstanceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(memcacheinstancesResource, memcacheinstancesKind, c.ns, opts), &v1beta1.MemcacheInstanceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.MemcacheInstanceList{ListMeta: obj.(*v1beta1.MemcacheInstanceList).ListMeta}
	for _, item := range obj.(*v1beta1.MemcacheInstanceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested memcacheInstances.
func (c *FakeMemcacheInstances) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(memcacheinstancesResource, c.ns, opts))

}

// Create takes the representation of a memcacheInstance and creates it.  Returns the server's representation of the memcacheInstance, and an error, if there is any.
func (c *FakeMemcacheInstances) Create(ctx context.Context, memcacheInstance *v1beta1.MemcacheInstance, opts v1.CreateOptions) (result *v1beta1.MemcacheInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(memcacheinstancesResource, c.ns, memcacheInstance), &v1beta1.MemcacheInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MemcacheInstance), err
}

// Update takes the representation of a memcacheInstance and updates it. Returns the server's representation of the memcacheInstance, and an error, if there is any.
func (c *FakeMemcacheInstances) Update(ctx context.Context, memcacheInstance *v1beta1.MemcacheInstance, opts v1.UpdateOptions) (result *v1beta1.MemcacheInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(memcacheinstancesResource, c.ns, memcacheInstance), &v1beta1.MemcacheInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MemcacheInstance), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeMemcacheInstances) UpdateStatus(ctx context.Context, memcacheInstance *v1beta1.MemcacheInstance, opts v1.UpdateOptions) (*v1beta1.MemcacheInstance, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(memcacheinstancesResource, "status", c.ns, memcacheInstance), &v1beta1.MemcacheInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MemcacheInstance), err
}

// Delete takes name of the memcacheInstance and deletes it. Returns an error if one occurs.
func (c *FakeMemcacheInstances) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(memcacheinstancesResource, c.ns, name, opts), &v1beta1.MemcacheInstance{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMemcacheInstances) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(memcacheinstancesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.MemcacheInstanceList{})
	return err
}

// Patch applies the patch and returns the patched memcacheInstance.
func (c *FakeMemcacheInstances) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.MemcacheInstance, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(memcacheinstancesResource, c.ns, name, pt, data, subresources...), &v1beta1.MemcacheInstance{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.MemcacheInstance), err
}
