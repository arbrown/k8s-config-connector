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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/binaryauthorization/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeBinaryAuthorizationPolicies implements BinaryAuthorizationPolicyInterface
type FakeBinaryAuthorizationPolicies struct {
	Fake *FakeBinaryauthorizationV1beta1
	ns   string
}

var binaryauthorizationpoliciesResource = schema.GroupVersionResource{Group: "binaryauthorization.cnrm.cloud.google.com", Version: "v1beta1", Resource: "binaryauthorizationpolicies"}

var binaryauthorizationpoliciesKind = schema.GroupVersionKind{Group: "binaryauthorization.cnrm.cloud.google.com", Version: "v1beta1", Kind: "BinaryAuthorizationPolicy"}

// Get takes name of the binaryAuthorizationPolicy, and returns the corresponding binaryAuthorizationPolicy object, and an error if there is any.
func (c *FakeBinaryAuthorizationPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(binaryauthorizationpoliciesResource, c.ns, name), &v1beta1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BinaryAuthorizationPolicy), err
}

// List takes label and field selectors, and returns the list of BinaryAuthorizationPolicies that match those selectors.
func (c *FakeBinaryAuthorizationPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.BinaryAuthorizationPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(binaryauthorizationpoliciesResource, binaryauthorizationpoliciesKind, c.ns, opts), &v1beta1.BinaryAuthorizationPolicyList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.BinaryAuthorizationPolicyList{ListMeta: obj.(*v1beta1.BinaryAuthorizationPolicyList).ListMeta}
	for _, item := range obj.(*v1beta1.BinaryAuthorizationPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested binaryAuthorizationPolicies.
func (c *FakeBinaryAuthorizationPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(binaryauthorizationpoliciesResource, c.ns, opts))

}

// Create takes the representation of a binaryAuthorizationPolicy and creates it.  Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationPolicies) Create(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.CreateOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(binaryauthorizationpoliciesResource, c.ns, binaryAuthorizationPolicy), &v1beta1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BinaryAuthorizationPolicy), err
}

// Update takes the representation of a binaryAuthorizationPolicy and updates it. Returns the server's representation of the binaryAuthorizationPolicy, and an error, if there is any.
func (c *FakeBinaryAuthorizationPolicies) Update(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(binaryauthorizationpoliciesResource, c.ns, binaryAuthorizationPolicy), &v1beta1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BinaryAuthorizationPolicy), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeBinaryAuthorizationPolicies) UpdateStatus(ctx context.Context, binaryAuthorizationPolicy *v1beta1.BinaryAuthorizationPolicy, opts v1.UpdateOptions) (*v1beta1.BinaryAuthorizationPolicy, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(binaryauthorizationpoliciesResource, "status", c.ns, binaryAuthorizationPolicy), &v1beta1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BinaryAuthorizationPolicy), err
}

// Delete takes name of the binaryAuthorizationPolicy and deletes it. Returns an error if one occurs.
func (c *FakeBinaryAuthorizationPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(binaryauthorizationpoliciesResource, c.ns, name, opts), &v1beta1.BinaryAuthorizationPolicy{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeBinaryAuthorizationPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(binaryauthorizationpoliciesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.BinaryAuthorizationPolicyList{})
	return err
}

// Patch applies the patch and returns the patched binaryAuthorizationPolicy.
func (c *FakeBinaryAuthorizationPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.BinaryAuthorizationPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(binaryauthorizationpoliciesResource, c.ns, name, pt, data, subresources...), &v1beta1.BinaryAuthorizationPolicy{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.BinaryAuthorizationPolicy), err
}
