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

	v1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/dialogflow/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeDialogflowIntents implements DialogflowIntentInterface
type FakeDialogflowIntents struct {
	Fake *FakeDialogflowV1alpha1
	ns   string
}

var dialogflowintentsResource = schema.GroupVersionResource{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "dialogflowintents"}

var dialogflowintentsKind = schema.GroupVersionKind{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowIntent"}

// Get takes name of the dialogflowIntent, and returns the corresponding dialogflowIntent object, and an error if there is any.
func (c *FakeDialogflowIntents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DialogflowIntent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dialogflowintentsResource, c.ns, name), &v1alpha1.DialogflowIntent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowIntent), err
}

// List takes label and field selectors, and returns the list of DialogflowIntents that match those selectors.
func (c *FakeDialogflowIntents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DialogflowIntentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dialogflowintentsResource, dialogflowintentsKind, c.ns, opts), &v1alpha1.DialogflowIntentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DialogflowIntentList{ListMeta: obj.(*v1alpha1.DialogflowIntentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DialogflowIntentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dialogflowIntents.
func (c *FakeDialogflowIntents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dialogflowintentsResource, c.ns, opts))

}

// Create takes the representation of a dialogflowIntent and creates it.  Returns the server's representation of the dialogflowIntent, and an error, if there is any.
func (c *FakeDialogflowIntents) Create(ctx context.Context, dialogflowIntent *v1alpha1.DialogflowIntent, opts v1.CreateOptions) (result *v1alpha1.DialogflowIntent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dialogflowintentsResource, c.ns, dialogflowIntent), &v1alpha1.DialogflowIntent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowIntent), err
}

// Update takes the representation of a dialogflowIntent and updates it. Returns the server's representation of the dialogflowIntent, and an error, if there is any.
func (c *FakeDialogflowIntents) Update(ctx context.Context, dialogflowIntent *v1alpha1.DialogflowIntent, opts v1.UpdateOptions) (result *v1alpha1.DialogflowIntent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dialogflowintentsResource, c.ns, dialogflowIntent), &v1alpha1.DialogflowIntent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowIntent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDialogflowIntents) UpdateStatus(ctx context.Context, dialogflowIntent *v1alpha1.DialogflowIntent, opts v1.UpdateOptions) (*v1alpha1.DialogflowIntent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dialogflowintentsResource, "status", c.ns, dialogflowIntent), &v1alpha1.DialogflowIntent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowIntent), err
}

// Delete takes name of the dialogflowIntent and deletes it. Returns an error if one occurs.
func (c *FakeDialogflowIntents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dialogflowintentsResource, c.ns, name, opts), &v1alpha1.DialogflowIntent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDialogflowIntents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dialogflowintentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DialogflowIntentList{})
	return err
}

// Patch applies the patch and returns the patched dialogflowIntent.
func (c *FakeDialogflowIntents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DialogflowIntent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dialogflowintentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DialogflowIntent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowIntent), err
}
