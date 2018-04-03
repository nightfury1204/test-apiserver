package fake

import (
	v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEyes implements EyeInterface
type FakeEyes struct {
	Fake *FakeTryapiV1alpha1
	ns   string
}

var eyesResource = schema.GroupVersionResource{Group: "tryapi.nahid.try", Version: "v1alpha1", Resource: "eyes"}

var eyesKind = schema.GroupVersionKind{Group: "tryapi.nahid.try", Version: "v1alpha1", Kind: "Eye"}

// Get takes name of the eye, and returns the corresponding eye object, and an error if there is any.
func (c *FakeEyes) Get(name string, options v1.GetOptions) (result *v1alpha1.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eyesResource, c.ns, name), &v1alpha1.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eye), err
}

// List takes label and field selectors, and returns the list of Eyes that match those selectors.
func (c *FakeEyes) List(opts v1.ListOptions) (result *v1alpha1.EyeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eyesResource, eyesKind, c.ns, opts), &v1alpha1.EyeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.EyeList{}
	for _, item := range obj.(*v1alpha1.EyeList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested eyes.
func (c *FakeEyes) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(eyesResource, c.ns, opts))

}

// Create takes the representation of a eye and creates it.  Returns the server's representation of the eye, and an error, if there is any.
func (c *FakeEyes) Create(eye *v1alpha1.Eye) (result *v1alpha1.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eyesResource, c.ns, eye), &v1alpha1.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eye), err
}

// Update takes the representation of a eye and updates it. Returns the server's representation of the eye, and an error, if there is any.
func (c *FakeEyes) Update(eye *v1alpha1.Eye) (result *v1alpha1.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eyesResource, c.ns, eye), &v1alpha1.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eye), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEyes) UpdateStatus(eye *v1alpha1.Eye) (*v1alpha1.Eye, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eyesResource, "status", c.ns, eye), &v1alpha1.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eye), err
}

// Delete takes name of the eye and deletes it. Returns an error if one occurs.
func (c *FakeEyes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eyesResource, c.ns, name), &v1alpha1.Eye{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEyes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eyesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.EyeList{})
	return err
}

// Patch applies the patch and returns the patched eye.
func (c *FakeEyes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eyesResource, c.ns, name, data, subresources...), &v1alpha1.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Eye), err
}
