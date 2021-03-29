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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "kubesphere.io/kubesphere/pkg/apis/application/v1alpha1"
	scheme "kubesphere.io/kubesphere/pkg/client/clientset/versioned/scheme"
)

// HelmReleasesGetter has a method to return a HelmReleaseInterface.
// A group's client should implement this interface.
type HelmReleasesGetter interface {
	HelmReleases() HelmReleaseInterface
}

// HelmReleaseInterface has methods to work with HelmRelease resources.
type HelmReleaseInterface interface {
	Create(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.CreateOptions) (*v1alpha1.HelmRelease, error)
	Update(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.UpdateOptions) (*v1alpha1.HelmRelease, error)
	UpdateStatus(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.UpdateOptions) (*v1alpha1.HelmRelease, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.HelmRelease, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.HelmReleaseList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmRelease, err error)
	HelmReleaseExpansion
}

// helmReleases implements HelmReleaseInterface
type helmReleases struct {
	client rest.Interface
}

// newHelmReleases returns a HelmReleases
func newHelmReleases(c *ApplicationV1alpha1Client) *helmReleases {
	return &helmReleases{
		client: c.RESTClient(),
	}
}

// Get takes name of the helmRelease, and returns the corresponding helmRelease object, and an error if there is any.
func (c *helmReleases) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.HelmRelease, err error) {
	result = &v1alpha1.HelmRelease{}
	err = c.client.Get().
		Resource("helmreleases").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of HelmReleases that match those selectors.
func (c *helmReleases) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.HelmReleaseList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.HelmReleaseList{}
	err = c.client.Get().
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested helmReleases.
func (c *helmReleases) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a helmRelease and creates it.  Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *helmReleases) Create(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.CreateOptions) (result *v1alpha1.HelmRelease, err error) {
	result = &v1alpha1.HelmRelease{}
	err = c.client.Post().
		Resource("helmreleases").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a helmRelease and updates it. Returns the server's representation of the helmRelease, and an error, if there is any.
func (c *helmReleases) Update(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.UpdateOptions) (result *v1alpha1.HelmRelease, err error) {
	result = &v1alpha1.HelmRelease{}
	err = c.client.Put().
		Resource("helmreleases").
		Name(helmRelease.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *helmReleases) UpdateStatus(ctx context.Context, helmRelease *v1alpha1.HelmRelease, opts v1.UpdateOptions) (result *v1alpha1.HelmRelease, err error) {
	result = &v1alpha1.HelmRelease{}
	err = c.client.Put().
		Resource("helmreleases").
		Name(helmRelease.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(helmRelease).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the helmRelease and deletes it. Returns an error if one occurs.
func (c *helmReleases) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("helmreleases").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *helmReleases) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Resource("helmreleases").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched helmRelease.
func (c *helmReleases) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.HelmRelease, err error) {
	result = &v1alpha1.HelmRelease{}
	err = c.client.Patch(pt).
		Resource("helmreleases").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
