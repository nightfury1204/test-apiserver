


package eye_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/client-go/rest"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/test"

	"github.com/nightfury1204/test-apiserver/pkg/apis"
	"github.com/nightfury1204/test-apiserver/pkg/client/clientset_generated/clientset"
	"github.com/nightfury1204/test-apiserver/pkg/openapi"
	"github.com/nightfury1204/test-apiserver/pkg/controller/sharedinformers"
	"github.com/nightfury1204/test-apiserver/pkg/controller/eye"
)

var testenv *test.TestEnvironment
var config *rest.Config
var cs *clientset.Clientset
var shutdown chan struct{}
var controller *eye.EyeController
var si *sharedinformers.SharedInformers

func TestEye(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "Eye Suite", []Reporter{test.NewlineReporter{}})
}

var _ = BeforeSuite(func() {
	testenv = test.NewTestEnvironment()
	config = testenv.Start(apis.GetAllApiBuilders(), openapi.GetOpenAPIDefinitions)
	cs = clientset.NewForConfigOrDie(config)

	shutdown = make(chan struct{})
	si = sharedinformers.NewSharedInformers(config, shutdown)
	controller = eye.NewEyeController(config, si)
	controller.Run(shutdown)
})

var _ = AfterSuite(func() {
	close(shutdown)
	testenv.Stop()
})
