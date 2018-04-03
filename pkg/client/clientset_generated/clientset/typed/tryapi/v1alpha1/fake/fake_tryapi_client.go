package fake

import (
	v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/clientset/typed/tryapi/v1alpha1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTryapiV1alpha1 struct {
	*testing.Fake
}

func (c *FakeTryapiV1alpha1) Eyes(namespace string) v1alpha1.EyeInterface {
	return &FakeEyes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTryapiV1alpha1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
