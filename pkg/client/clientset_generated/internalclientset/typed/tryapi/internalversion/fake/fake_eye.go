package fake

import (
	tryapi "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeEyes implements EyeInterface
type FakeEyes struct {
	Fake *FakeTryapi
	ns   string
}

var eyesResource = schema.GroupVersionResource{Group: "tryapi.nahid.try", Version: "", Resource: "eyes"}

var eyesKind = schema.GroupVersionKind{Group: "tryapi.nahid.try", Version: "", Kind: "Eye"}

// Get takes name of the eye, and returns the corresponding eye object, and an error if there is any.
func (c *FakeEyes) Get(name string, options v1.GetOptions) (result *tryapi.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(eyesResource, c.ns, name), &tryapi.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tryapi.Eye), err
}

// List takes label and field selectors, and returns the list of Eyes that match those selectors.
func (c *FakeEyes) List(opts v1.ListOptions) (result *tryapi.EyeList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(eyesResource, eyesKind, c.ns, opts), &tryapi.EyeList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &tryapi.EyeList{}
	for _, item := range obj.(*tryapi.EyeList).Items {
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
func (c *FakeEyes) Create(eye *tryapi.Eye) (result *tryapi.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(eyesResource, c.ns, eye), &tryapi.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tryapi.Eye), err
}

// Update takes the representation of a eye and updates it. Returns the server's representation of the eye, and an error, if there is any.
func (c *FakeEyes) Update(eye *tryapi.Eye) (result *tryapi.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(eyesResource, c.ns, eye), &tryapi.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tryapi.Eye), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeEyes) UpdateStatus(eye *tryapi.Eye) (*tryapi.Eye, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(eyesResource, "status", c.ns, eye), &tryapi.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tryapi.Eye), err
}

// Delete takes name of the eye and deletes it. Returns an error if one occurs.
func (c *FakeEyes) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(eyesResource, c.ns, name), &tryapi.Eye{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeEyes) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(eyesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &tryapi.EyeList{})
	return err
}

// Patch applies the patch and returns the patched eye.
func (c *FakeEyes) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *tryapi.Eye, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(eyesResource, c.ns, name, data, subresources...), &tryapi.Eye{})

	if obj == nil {
		return nil, err
	}
	return obj.(*tryapi.Eye), err
}
