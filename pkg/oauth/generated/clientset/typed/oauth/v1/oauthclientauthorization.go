package v1

import (
	v1 "github.com/openshift/origin/pkg/oauth/apis/oauth/v1"
	scheme "github.com/openshift/origin/pkg/oauth/generated/clientset/scheme"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OAuthClientAuthorizationsGetter has a method to return a OAuthClientAuthorizationInterface.
// A group's client should implement this interface.
type OAuthClientAuthorizationsGetter interface {
	OAuthClientAuthorizations() OAuthClientAuthorizationInterface
}

// OAuthClientAuthorizationInterface has methods to work with OAuthClientAuthorization resources.
type OAuthClientAuthorizationInterface interface {
	Create(*v1.OAuthClientAuthorization) (*v1.OAuthClientAuthorization, error)
	Update(*v1.OAuthClientAuthorization) (*v1.OAuthClientAuthorization, error)
	Delete(name string, options *meta_v1.DeleteOptions) error
	DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error
	Get(name string, options meta_v1.GetOptions) (*v1.OAuthClientAuthorization, error)
	List(opts meta_v1.ListOptions) (*v1.OAuthClientAuthorizationList, error)
	Watch(opts meta_v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OAuthClientAuthorization, err error)
	OAuthClientAuthorizationExpansion
}

// oAuthClientAuthorizations implements OAuthClientAuthorizationInterface
type oAuthClientAuthorizations struct {
	client rest.Interface
}

// newOAuthClientAuthorizations returns a OAuthClientAuthorizations
func newOAuthClientAuthorizations(c *OauthV1Client) *oAuthClientAuthorizations {
	return &oAuthClientAuthorizations{
		client: c.RESTClient(),
	}
}

// Get takes name of the oAuthClientAuthorization, and returns the corresponding oAuthClientAuthorization object, and an error if there is any.
func (c *oAuthClientAuthorizations) Get(name string, options meta_v1.GetOptions) (result *v1.OAuthClientAuthorization, err error) {
	result = &v1.OAuthClientAuthorization{}
	err = c.client.Get().
		Resource("oauthclientauthorizations").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OAuthClientAuthorizations that match those selectors.
func (c *oAuthClientAuthorizations) List(opts meta_v1.ListOptions) (result *v1.OAuthClientAuthorizationList, err error) {
	result = &v1.OAuthClientAuthorizationList{}
	err = c.client.Get().
		Resource("oauthclientauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested oAuthClientAuthorizations.
func (c *oAuthClientAuthorizations) Watch(opts meta_v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Resource("oauthclientauthorizations").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a oAuthClientAuthorization and creates it.  Returns the server's representation of the oAuthClientAuthorization, and an error, if there is any.
func (c *oAuthClientAuthorizations) Create(oAuthClientAuthorization *v1.OAuthClientAuthorization) (result *v1.OAuthClientAuthorization, err error) {
	result = &v1.OAuthClientAuthorization{}
	err = c.client.Post().
		Resource("oauthclientauthorizations").
		Body(oAuthClientAuthorization).
		Do().
		Into(result)
	return
}

// Update takes the representation of a oAuthClientAuthorization and updates it. Returns the server's representation of the oAuthClientAuthorization, and an error, if there is any.
func (c *oAuthClientAuthorizations) Update(oAuthClientAuthorization *v1.OAuthClientAuthorization) (result *v1.OAuthClientAuthorization, err error) {
	result = &v1.OAuthClientAuthorization{}
	err = c.client.Put().
		Resource("oauthclientauthorizations").
		Name(oAuthClientAuthorization.Name).
		Body(oAuthClientAuthorization).
		Do().
		Into(result)
	return
}

// Delete takes name of the oAuthClientAuthorization and deletes it. Returns an error if one occurs.
func (c *oAuthClientAuthorizations) Delete(name string, options *meta_v1.DeleteOptions) error {
	return c.client.Delete().
		Resource("oauthclientauthorizations").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *oAuthClientAuthorizations) DeleteCollection(options *meta_v1.DeleteOptions, listOptions meta_v1.ListOptions) error {
	return c.client.Delete().
		Resource("oauthclientauthorizations").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched oAuthClientAuthorization.
func (c *oAuthClientAuthorizations) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1.OAuthClientAuthorization, err error) {
	result = &v1.OAuthClientAuthorization{}
	err = c.client.Patch(pt).
		Resource("oauthclientauthorizations").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
