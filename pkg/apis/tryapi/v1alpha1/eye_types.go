


package v1alpha1

import (
	"log"

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apiserver/pkg/endpoints/request"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/nightfury1204/test-apiserver/pkg/apis/tryapi"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Eye
// +k8s:openapi-gen=true
// +resource:path=eyes,strategy=EyeStrategy
type Eye struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   EyeSpec   `json:"spec,omitempty"`
	Status EyeStatus `json:"status,omitempty"`
}

// EyeSpec defines the desired state of Eye
type EyeSpec struct {
	Description string `json:"description",omitempty`
}

// EyeStatus defines the observed state of Eye
type EyeStatus struct {
	Ok bool `json:"ok"`
}

// Validate checks that an instance of Eye is well formed
func (EyeStrategy) Validate(ctx request.Context, obj runtime.Object) field.ErrorList {
	o := obj.(*tryapi.Eye)
	log.Printf("Validating fields for Eye %s\n", o.Name)
	errors := field.ErrorList{}
	// perform validation here and add to errors using field.Invalid
	if o.Spec.Description == "" {
		errors = append(errors, field.Invalid(field.NewPath("spec","description"),o.Spec.Description,"must be non-empty"))
	}
	return errors
}

// DefaultingFunction sets default Eye field values
func (EyeSchemeFns) DefaultingFunction(o interface{}) {
	obj := o.(*Eye)
	// set default field values here
	log.Printf("Defaulting fields for Eye %s\n", obj.Name)
}
