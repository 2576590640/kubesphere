/*
Copyright 2020 The KubeSphere Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubesphere.io/kubesphere/pkg/apis/types/v1beta1"
)

// FakeFederatedWorkspaces implements FederatedWorkspaceInterface
type FakeFederatedWorkspaces struct {
	Fake *FakeTypesV1beta1
}

var federatedworkspacesResource = schema.GroupVersionResource{Group: "types.kubefed.io", Version: "v1beta1", Resource: "federatedworkspaces"}

var federatedworkspacesKind = schema.GroupVersionKind{Group: "types.kubefed.io", Version: "v1beta1", Kind: "FederatedWorkspace"}

// Get takes name of the federatedWorkspace, and returns the corresponding federatedWorkspace object, and an error if there is any.
func (c *FakeFederatedWorkspaces) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.FederatedWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(federatedworkspacesResource, name), &v1beta1.FederatedWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedWorkspace), err
}

// List takes label and field selectors, and returns the list of FederatedWorkspaces that match those selectors.
func (c *FakeFederatedWorkspaces) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.FederatedWorkspaceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(federatedworkspacesResource, federatedworkspacesKind, opts), &v1beta1.FederatedWorkspaceList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.FederatedWorkspaceList{ListMeta: obj.(*v1beta1.FederatedWorkspaceList).ListMeta}
	for _, item := range obj.(*v1beta1.FederatedWorkspaceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested federatedWorkspaces.
func (c *FakeFederatedWorkspaces) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(federatedworkspacesResource, opts))
}

// Create takes the representation of a federatedWorkspace and creates it.  Returns the server's representation of the federatedWorkspace, and an error, if there is any.
func (c *FakeFederatedWorkspaces) Create(ctx context.Context, federatedWorkspace *v1beta1.FederatedWorkspace, opts v1.CreateOptions) (result *v1beta1.FederatedWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(federatedworkspacesResource, federatedWorkspace), &v1beta1.FederatedWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedWorkspace), err
}

// Update takes the representation of a federatedWorkspace and updates it. Returns the server's representation of the federatedWorkspace, and an error, if there is any.
func (c *FakeFederatedWorkspaces) Update(ctx context.Context, federatedWorkspace *v1beta1.FederatedWorkspace, opts v1.UpdateOptions) (result *v1beta1.FederatedWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(federatedworkspacesResource, federatedWorkspace), &v1beta1.FederatedWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedWorkspace), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeFederatedWorkspaces) UpdateStatus(ctx context.Context, federatedWorkspace *v1beta1.FederatedWorkspace, opts v1.UpdateOptions) (*v1beta1.FederatedWorkspace, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(federatedworkspacesResource, "status", federatedWorkspace), &v1beta1.FederatedWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedWorkspace), err
}

// Delete takes name of the federatedWorkspace and deletes it. Returns an error if one occurs.
func (c *FakeFederatedWorkspaces) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(federatedworkspacesResource, name), &v1beta1.FederatedWorkspace{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFederatedWorkspaces) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(federatedworkspacesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.FederatedWorkspaceList{})
	return err
}

// Patch applies the patch and returns the patched federatedWorkspace.
func (c *FakeFederatedWorkspaces) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.FederatedWorkspace, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(federatedworkspacesResource, name, pt, data, subresources...), &v1beta1.FederatedWorkspace{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.FederatedWorkspace), err
}
