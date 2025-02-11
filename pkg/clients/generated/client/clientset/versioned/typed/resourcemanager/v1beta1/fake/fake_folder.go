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

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/resourcemanager/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFolders implements FolderInterface
type FakeFolders struct {
	Fake *FakeResourcemanagerV1beta1
	ns   string
}

var foldersResource = schema.GroupVersionResource{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Resource: "folders"}

var foldersKind = schema.GroupVersionKind{Group: "resourcemanager.cnrm.cloud.google.com", Version: "v1beta1", Kind: "Folder"}

// Get takes name of the folder, and returns the corresponding folder object, and an error if there is any.
func (c *FakeFolders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(foldersResource, c.ns, name), &v1beta1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Folder), err
}

// List takes label and field selectors, and returns the list of Folders that match those selectors.
func (c *FakeFolders) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FolderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(foldersResource, foldersKind, c.ns, opts), &v1beta1.FolderList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FolderList{ListMeta: obj.(*v1beta1.FolderList).ListMeta}
	for _, item := range obj.(*v1beta1.FolderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested folders.
func (c *FakeFolders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(foldersResource, c.ns, opts))

}

// Create takes the representation of a folder and creates it.  Returns the server's representation of the folder, and an error, if there is any.
func (c *FakeFolders) Create(ctx context.Context, folder *v1beta1.Folder, opts v1.CreateOptions) (result *v1beta1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(foldersResource, c.ns, folder), &v1beta1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Folder), err
}

// Update takes the representation of a folder and updates it. Returns the server's representation of the folder, and an error, if there is any.
func (c *FakeFolders) Update(ctx context.Context, folder *v1beta1.Folder, opts v1.UpdateOptions) (result *v1beta1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(foldersResource, c.ns, folder), &v1beta1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Folder), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFolders) UpdateStatus(ctx context.Context, folder *v1beta1.Folder, opts v1.UpdateOptions) (*v1beta1.Folder, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(foldersResource, "status", c.ns, folder), &v1beta1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Folder), err
}

// Delete takes name of the folder and deletes it. Returns an error if one occurs.
func (c *FakeFolders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(foldersResource, c.ns, name, opts), &v1beta1.Folder{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFolders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(foldersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FolderList{})
	return err
}

// Patch applies the patch and returns the patched folder.
func (c *FakeFolders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.Folder, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(foldersResource, c.ns, name, pt, data, subresources...), &v1beta1.Folder{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.Folder), err
}
