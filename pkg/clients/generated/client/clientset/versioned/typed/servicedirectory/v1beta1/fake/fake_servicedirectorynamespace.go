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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/servicedirectory/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServiceDirectoryNamespaces implements ServiceDirectoryNamespaceInterface
type FakeServiceDirectoryNamespaces struct {
	Fake *FakeServicedirectoryV1beta1
	ns   string
}

var servicedirectorynamespacesResource = schema.GroupVersionResource{Group: "servicedirectory.cnrm.cloud.google.com", Version: "v1beta1", Resource: "servicedirectorynamespaces"}

var servicedirectorynamespacesKind = schema.GroupVersionKind{Group: "servicedirectory.cnrm.cloud.google.com", Version: "v1beta1", Kind: "ServiceDirectoryNamespace"}

// Get takes name of the serviceDirectoryNamespace, and returns the corresponding serviceDirectoryNamespace object, and an error if there is any.
func (c *FakeServiceDirectoryNamespaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ServiceDirectoryNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servicedirectorynamespacesResource, c.ns, name), &v1beta1.ServiceDirectoryNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceDirectoryNamespace), err
}

// List takes label and field selectors, and returns the list of ServiceDirectoryNamespaces that match those selectors.
func (c *FakeServiceDirectoryNamespaces) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ServiceDirectoryNamespaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servicedirectorynamespacesResource, servicedirectorynamespacesKind, c.ns, opts), &v1beta1.ServiceDirectoryNamespaceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ServiceDirectoryNamespaceList{ListMeta: obj.(*v1beta1.ServiceDirectoryNamespaceList).ListMeta}
	for _, item := range obj.(*v1beta1.ServiceDirectoryNamespaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested serviceDirectoryNamespaces.
func (c *FakeServiceDirectoryNamespaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servicedirectorynamespacesResource, c.ns, opts))

}

// Create takes the representation of a serviceDirectoryNamespace and creates it.  Returns the server's representation of the serviceDirectoryNamespace, and an error, if there is any.
func (c *FakeServiceDirectoryNamespaces) Create(ctx context.Context, serviceDirectoryNamespace *v1beta1.ServiceDirectoryNamespace, opts v1.CreateOptions) (result *v1beta1.ServiceDirectoryNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servicedirectorynamespacesResource, c.ns, serviceDirectoryNamespace), &v1beta1.ServiceDirectoryNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceDirectoryNamespace), err
}

// Update takes the representation of a serviceDirectoryNamespace and updates it. Returns the server's representation of the serviceDirectoryNamespace, and an error, if there is any.
func (c *FakeServiceDirectoryNamespaces) Update(ctx context.Context, serviceDirectoryNamespace *v1beta1.ServiceDirectoryNamespace, opts v1.UpdateOptions) (result *v1beta1.ServiceDirectoryNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servicedirectorynamespacesResource, c.ns, serviceDirectoryNamespace), &v1beta1.ServiceDirectoryNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceDirectoryNamespace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServiceDirectoryNamespaces) UpdateStatus(ctx context.Context, serviceDirectoryNamespace *v1beta1.ServiceDirectoryNamespace, opts v1.UpdateOptions) (*v1beta1.ServiceDirectoryNamespace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servicedirectorynamespacesResource, "status", c.ns, serviceDirectoryNamespace), &v1beta1.ServiceDirectoryNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceDirectoryNamespace), err
}

// Delete takes name of the serviceDirectoryNamespace and deletes it. Returns an error if one occurs.
func (c *FakeServiceDirectoryNamespaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servicedirectorynamespacesResource, c.ns, name, opts), &v1beta1.ServiceDirectoryNamespace{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServiceDirectoryNamespaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servicedirectorynamespacesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ServiceDirectoryNamespaceList{})
	return err
}

// Patch applies the patch and returns the patched serviceDirectoryNamespace.
func (c *FakeServiceDirectoryNamespaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ServiceDirectoryNamespace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servicedirectorynamespacesResource, c.ns, name, pt, data, subresources...), &v1beta1.ServiceDirectoryNamespace{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ServiceDirectoryNamespace), err
}
