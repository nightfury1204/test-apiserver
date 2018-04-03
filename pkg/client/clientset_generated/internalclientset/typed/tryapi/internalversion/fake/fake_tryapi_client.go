package fake

import (
	internalversion "github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/internalclientset/typed/tryapi/internalversion"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeTryapi struct {
	*testing.Fake
}

func (c *FakeTryapi) Eyes(namespace string) internalversion.EyeInterface {
	return &FakeEyes{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeTryapi) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
