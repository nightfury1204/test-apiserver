package internalversion

import (
	"github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/internalclientset/scheme"
	rest "k8s.io/client-go/rest"
)

type TryapiInterface interface {
	RESTClient() rest.Interface
	EyesGetter
}

// TryapiClient is used to interact with features provided by the tryapi.nahid.try group.
type TryapiClient struct {
	restClient rest.Interface
}

func (c *TryapiClient) Eyes(namespace string) EyeInterface {
	return newEyes(c, namespace)
}

// NewForConfig creates a new TryapiClient for the given config.
func NewForConfig(c *rest.Config) (*TryapiClient, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TryapiClient{client}, nil
}

// NewForConfigOrDie creates a new TryapiClient for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TryapiClient {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TryapiClient for the given RESTClient.
func New(c rest.Interface) *TryapiClient {
	return &TryapiClient{c}
}

func setConfigDefaults(config *rest.Config) error {
	g, err := scheme.Registry.Group("tryapi.nahid.try")
	if err != nil {
		return err
	}

	config.APIPath = "/apis"
	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}
	if config.GroupVersion == nil || config.GroupVersion.Group != g.GroupVersion.Group {
		gv := g.GroupVersion
		config.GroupVersion = &gv
	}
	config.NegotiatedSerializer = scheme.Codecs

	if config.QPS == 0 {
		config.QPS = 5
	}
	if config.Burst == 0 {
		config.Burst = 10
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TryapiClient) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
