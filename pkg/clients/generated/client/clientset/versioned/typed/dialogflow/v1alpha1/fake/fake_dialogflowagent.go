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

// FakeDialogflowAgents implements DialogflowAgentInterface
type FakeDialogflowAgents struct {
	Fake *FakeDialogflowV1alpha1
	ns   string
}

var dialogflowagentsResource = schema.GroupVersionResource{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Resource: "dialogflowagents"}

var dialogflowagentsKind = schema.GroupVersionKind{Group: "dialogflow.cnrm.cloud.google.com", Version: "v1alpha1", Kind: "DialogflowAgent"}

// Get takes name of the dialogflowAgent, and returns the corresponding dialogflowAgent object, and an error if there is any.
func (c *FakeDialogflowAgents) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.DialogflowAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(dialogflowagentsResource, c.ns, name), &v1alpha1.DialogflowAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowAgent), err
}

// List takes label and field selectors, and returns the list of DialogflowAgents that match those selectors.
func (c *FakeDialogflowAgents) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.DialogflowAgentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(dialogflowagentsResource, dialogflowagentsKind, c.ns, opts), &v1alpha1.DialogflowAgentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.DialogflowAgentList{ListMeta: obj.(*v1alpha1.DialogflowAgentList).ListMeta}
	for _, item := range obj.(*v1alpha1.DialogflowAgentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested dialogflowAgents.
func (c *FakeDialogflowAgents) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(dialogflowagentsResource, c.ns, opts))

}

// Create takes the representation of a dialogflowAgent and creates it.  Returns the server's representation of the dialogflowAgent, and an error, if there is any.
func (c *FakeDialogflowAgents) Create(ctx context.Context, dialogflowAgent *v1alpha1.DialogflowAgent, opts v1.CreateOptions) (result *v1alpha1.DialogflowAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(dialogflowagentsResource, c.ns, dialogflowAgent), &v1alpha1.DialogflowAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowAgent), err
}

// Update takes the representation of a dialogflowAgent and updates it. Returns the server's representation of the dialogflowAgent, and an error, if there is any.
func (c *FakeDialogflowAgents) Update(ctx context.Context, dialogflowAgent *v1alpha1.DialogflowAgent, opts v1.UpdateOptions) (result *v1alpha1.DialogflowAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(dialogflowagentsResource, c.ns, dialogflowAgent), &v1alpha1.DialogflowAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowAgent), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeDialogflowAgents) UpdateStatus(ctx context.Context, dialogflowAgent *v1alpha1.DialogflowAgent, opts v1.UpdateOptions) (*v1alpha1.DialogflowAgent, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(dialogflowagentsResource, "status", c.ns, dialogflowAgent), &v1alpha1.DialogflowAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowAgent), err
}

// Delete takes name of the dialogflowAgent and deletes it. Returns an error if one occurs.
func (c *FakeDialogflowAgents) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(dialogflowagentsResource, c.ns, name, opts), &v1alpha1.DialogflowAgent{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeDialogflowAgents) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(dialogflowagentsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.DialogflowAgentList{})
	return err
}

// Patch applies the patch and returns the patched dialogflowAgent.
func (c *FakeDialogflowAgents) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.DialogflowAgent, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(dialogflowagentsResource, c.ns, name, pt, data, subresources...), &v1alpha1.DialogflowAgent{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.DialogflowAgent), err
}
