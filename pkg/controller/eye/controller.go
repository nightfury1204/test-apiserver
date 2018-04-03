


package eye

import (
	"log"

	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"

	"github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
	"github.com/nightfury1204/test-apiserver/pkg/controller/sharedinformers"
	listers "github.com/nightfury1204/test-apiserver/pkg/client/listers_generated/tryapi/v1alpha1"
)

// +controller:group=tryapi,version=v1alpha1,kind=Eye,resource=eyes
type EyeControllerImpl struct {
	builders.DefaultControllerFns

	// lister indexes properties about Eye
	lister listers.EyeLister
}

// Init initializes the controller and is called by the generated code
// Register watches for additional resource types here.
func (c *EyeControllerImpl) Init(arguments sharedinformers.ControllerInitArguments) {
	// Use the lister for indexing eyes labels
	c.lister = arguments.GetSharedInformers().Factory.Tryapi().V1alpha1().Eyes().Lister()
}

// Reconcile handles enqueued messages
func (c *EyeControllerImpl) Reconcile(u *v1alpha1.Eye) error {
	// Implement controller logic here
	log.Printf("Running reconcile Eye for %s\n", u.Name)
	return nil
}

func (c *EyeControllerImpl) Get(namespace, name string) (*v1alpha1.Eye, error) {
	return c.lister.Eyes(namespace).Get(name)
}
