package player

import (
	"fmt"

	"k8s.io/apimachinery/pkg/fields"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apiserver/pkg/registry/generic"
	"k8s.io/apiserver/pkg/storage"
	"k8s.io/apiserver/pkg/storage/names"

	genericapirequest "k8s.io/apiserver/pkg/endpoints/request"
	"github.com/nightfury1204/test-apiserver/pkg/apis/tryapi"
)

// NewStrategy creates and returns a playerStrategy instance
func NewStrategy(typer runtime.ObjectTyper) playerStrategy {
	return playerStrategy{typer, names.SimpleNameGenerator}
}

// GetAttrs returns labels.Set, fields.Set, the presence of Initializers if any
// and error in case the given runtime.Object is not a Flunder
func GetAttrs(obj runtime.Object) (labels.Set, fields.Set, bool, error) {
	apiserver, ok := obj.(*tryapi.Player)
	if !ok {
		return nil, nil, false, fmt.Errorf("given object is not a Flunder")
	}
	return labels.Set(apiserver.ObjectMeta.Labels), SelectableFields(apiserver), apiserver.Initializers != nil, nil
}

// MatchPlayer is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchPlayer(label labels.Selector, field fields.Selector) storage.SelectionPredicate {
	return storage.SelectionPredicate{
		Label:    label,
		Field:    field,
		GetAttrs: GetAttrs,
	}
}

// SelectableFields returns a field set that represents the object.
func SelectableFields(obj *tryapi.Player) fields.Set {
	return generic.ObjectMetaFieldsSet(&obj.ObjectMeta, true)
}

type playerStrategy struct {
	runtime.ObjectTyper
	names.NameGenerator
}

func (playerStrategy) NamespaceScoped() bool {
	return true
}

func (playerStrategy) PrepareForCreate(ctx genericapirequest.Context, obj runtime.Object) {
}

func (playerStrategy) PrepareForUpdate(ctx genericapirequest.Context, obj, old runtime.Object) {
}

func (playerStrategy) Validate(ctx genericapirequest.Context, obj runtime.Object) field.ErrorList {
	return field.ErrorList{}
}

func (playerStrategy) AllowCreateOnUpdate() bool {
	return false
}

func (playerStrategy) AllowUnconditionalUpdate() bool {
	return false
}

func (playerStrategy) Canonicalize(obj runtime.Object) {
}

func (playerStrategy) ValidateUpdate(ctx genericapirequest.Context, obj, old runtime.Object) field.ErrorList {
	return field.ErrorList{}
}