package v1alpha1

import (
	v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
	"github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/clientset/scheme"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	rest "k8s.io/client-go/rest"
)

type TryapiV1alpha1Interface interface {
	RESTClient() rest.Interface
	EyesGetter
}

// TryapiV1alpha1Client is used to interact with features provided by the tryapi.nahid.try group.
type TryapiV1alpha1Client struct {
	restClient rest.Interface
}

func (c *TryapiV1alpha1Client) Eyes(namespace string) EyeInterface {
	return newEyes(c, namespace)
}

// NewForConfig creates a new TryapiV1alpha1Client for the given config.
func NewForConfig(c *rest.Config) (*TryapiV1alpha1Client, error) {
	config := *c
	if err := setConfigDefaults(&config); err != nil {
		return nil, err
	}
	client, err := rest.RESTClientFor(&config)
	if err != nil {
		return nil, err
	}
	return &TryapiV1alpha1Client{client}, nil
}

// NewForConfigOrDie creates a new TryapiV1alpha1Client for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *TryapiV1alpha1Client {
	client, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return client
}

// New creates a new TryapiV1alpha1Client for the given RESTClient.
func New(c rest.Interface) *TryapiV1alpha1Client {
	return &TryapiV1alpha1Client{c}
}

func setConfigDefaults(config *rest.Config) error {
	gv := v1alpha1.SchemeGroupVersion
	config.GroupVersion = &gv
	config.APIPath = "/apis"
	config.NegotiatedSerializer = serializer.DirectCodecFactory{CodecFactory: scheme.Codecs}

	if config.UserAgent == "" {
		config.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	return nil
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *TryapiV1alpha1Client) RESTClient() rest.Interface {
	if c == nil {
		return nil
	}
	return c.restClient
}
