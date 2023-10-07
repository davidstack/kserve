/*
Copyright 2023 The KServe Authors.

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

	v1alpha1 "github.com/kserve/kserve/pkg/apis/serving/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeServingRuntimes implements ServingRuntimeInterface
type FakeServingRuntimes struct {
	Fake *FakeServingV1alpha1
	ns   string
}

var servingruntimesResource = v1alpha1.SchemeGroupVersion.WithResource("servingruntimes")

var servingruntimesKind = v1alpha1.SchemeGroupVersion.WithKind("ServingRuntime")

// Get takes name of the servingRuntime, and returns the corresponding servingRuntime object, and an error if there is any.
func (c *FakeServingRuntimes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ServingRuntime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(servingruntimesResource, c.ns, name), &v1alpha1.ServingRuntime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServingRuntime), err
}

// List takes label and field selectors, and returns the list of ServingRuntimes that match those selectors.
func (c *FakeServingRuntimes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ServingRuntimeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(servingruntimesResource, servingruntimesKind, c.ns, opts), &v1alpha1.ServingRuntimeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ServingRuntimeList{ListMeta: obj.(*v1alpha1.ServingRuntimeList).ListMeta}
	for _, item := range obj.(*v1alpha1.ServingRuntimeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested servingRuntimes.
func (c *FakeServingRuntimes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(servingruntimesResource, c.ns, opts))

}

// Create takes the representation of a servingRuntime and creates it.  Returns the server's representation of the servingRuntime, and an error, if there is any.
func (c *FakeServingRuntimes) Create(ctx context.Context, servingRuntime *v1alpha1.ServingRuntime, opts v1.CreateOptions) (result *v1alpha1.ServingRuntime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(servingruntimesResource, c.ns, servingRuntime), &v1alpha1.ServingRuntime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServingRuntime), err
}

// Update takes the representation of a servingRuntime and updates it. Returns the server's representation of the servingRuntime, and an error, if there is any.
func (c *FakeServingRuntimes) Update(ctx context.Context, servingRuntime *v1alpha1.ServingRuntime, opts v1.UpdateOptions) (result *v1alpha1.ServingRuntime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(servingruntimesResource, c.ns, servingRuntime), &v1alpha1.ServingRuntime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServingRuntime), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeServingRuntimes) UpdateStatus(ctx context.Context, servingRuntime *v1alpha1.ServingRuntime, opts v1.UpdateOptions) (*v1alpha1.ServingRuntime, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(servingruntimesResource, "status", c.ns, servingRuntime), &v1alpha1.ServingRuntime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServingRuntime), err
}

// Delete takes name of the servingRuntime and deletes it. Returns an error if one occurs.
func (c *FakeServingRuntimes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(servingruntimesResource, c.ns, name, opts), &v1alpha1.ServingRuntime{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeServingRuntimes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(servingruntimesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ServingRuntimeList{})
	return err
}

// Patch applies the patch and returns the patched servingRuntime.
func (c *FakeServingRuntimes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ServingRuntime, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(servingruntimesResource, c.ns, name, pt, data, subresources...), &v1alpha1.ServingRuntime{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ServingRuntime), err
}
