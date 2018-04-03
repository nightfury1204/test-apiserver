package internalversion

import (
	tryapi "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi"
	scheme "github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// EyesGetter has a method to return a EyeInterface.
// A group's client should implement this interface.
type EyesGetter interface {
	Eyes(namespace string) EyeInterface
}

// EyeInterface has methods to work with Eye resources.
type EyeInterface interface {
	Create(*tryapi.Eye) (*tryapi.Eye, error)
	Update(*tryapi.Eye) (*tryapi.Eye, error)
	UpdateStatus(*tryapi.Eye) (*tryapi.Eye, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*tryapi.Eye, error)
	List(opts v1.ListOptions) (*tryapi.EyeList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *tryapi.Eye, err error)
	EyeExpansion
}

// eyes implements EyeInterface
type eyes struct {
	client rest.Interface
	ns     string
}

// newEyes returns a Eyes
func newEyes(c *TryapiClient, namespace string) *eyes {
	return &eyes{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the eye, and returns the corresponding eye object, and an error if there is any.
func (c *eyes) Get(name string, options v1.GetOptions) (result *tryapi.Eye, err error) {
	result = &tryapi.Eye{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eyes").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Eyes that match those selectors.
func (c *eyes) List(opts v1.ListOptions) (result *tryapi.EyeList, err error) {
	result = &tryapi.EyeList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("eyes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested eyes.
func (c *eyes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("eyes").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a eye and creates it.  Returns the server's representation of the eye, and an error, if there is any.
func (c *eyes) Create(eye *tryapi.Eye) (result *tryapi.Eye, err error) {
	result = &tryapi.Eye{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("eyes").
		Body(eye).
		Do().
		Into(result)
	return
}

// Update takes the representation of a eye and updates it. Returns the server's representation of the eye, and an error, if there is any.
func (c *eyes) Update(eye *tryapi.Eye) (result *tryapi.Eye, err error) {
	result = &tryapi.Eye{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eyes").
		Name(eye.Name).
		Body(eye).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *eyes) UpdateStatus(eye *tryapi.Eye) (result *tryapi.Eye, err error) {
	result = &tryapi.Eye{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("eyes").
		Name(eye.Name).
		SubResource("status").
		Body(eye).
		Do().
		Into(result)
	return
}

// Delete takes name of the eye and deletes it. Returns an error if one occurs.
func (c *eyes) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eyes").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *eyes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("eyes").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched eye.
func (c *eyes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *tryapi.Eye, err error) {
	result = &tryapi.Eye{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("eyes").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
