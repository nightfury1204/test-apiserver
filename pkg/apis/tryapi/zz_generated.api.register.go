/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// This file was autogenerated by apiregister-gen. Do not edit it manually!

package tryapi

import (
	"fmt"
	"github.com/kubernetes-incubator/apiserver-builder/pkg/builders"
	"k8s.io/apimachinery/pkg/apis/meta/internalversion"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apiserver/pkg/endpoints/request"
	"k8s.io/apiserver/pkg/registry/rest"
)

var (
	InternalEye = builders.NewInternalResource(
		"eyes",
		"Eye",
		func() runtime.Object { return &Eye{} },
		func() runtime.Object { return &EyeList{} },
	)
	InternalEyeStatus = builders.NewInternalResourceStatus(
		"eyes",
		"EyeStatus",
		func() runtime.Object { return &Eye{} },
		func() runtime.Object { return &EyeList{} },
	)
	// Registered resources and subresources
	ApiVersion = builders.NewApiGroup("tryapi.nahid.try").WithKinds(
		InternalEye,
		InternalEyeStatus,
	)

	// Required by code generated by go2idl
	AddToScheme        = ApiVersion.SchemaBuilder.AddToScheme
	SchemeBuilder      = ApiVersion.SchemaBuilder
	localSchemeBuilder = &SchemeBuilder
	SchemeGroupVersion = ApiVersion.GroupVersion
)

// Required by code generated by go2idl
// Kind takes an unqualified kind and returns a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Required by code generated by go2idl
// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// +genclient
// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Eye struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Spec   EyeSpec
	Status EyeStatus
}

type EyeSpec struct {
	Description string
}

type EyeStatus struct {
	Ok bool
}

//
// Eye Functions and Structs
//
// +k8s:deepcopy-gen=false
type EyeStrategy struct {
	builders.DefaultStorageStrategy
}

// +k8s:deepcopy-gen=false
type EyeStatusStrategy struct {
	builders.DefaultStatusStorageStrategy
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EyeList struct {
	metav1.TypeMeta
	metav1.ListMeta
	Items []Eye
}

func (Eye) NewStatus() interface{} {
	return EyeStatus{}
}

func (pc *Eye) GetStatus() interface{} {
	return pc.Status
}

func (pc *Eye) SetStatus(s interface{}) {
	pc.Status = s.(EyeStatus)
}

func (pc *Eye) GetSpec() interface{} {
	return pc.Spec
}

func (pc *Eye) SetSpec(s interface{}) {
	pc.Spec = s.(EyeSpec)
}

func (pc *Eye) GetObjectMeta() *metav1.ObjectMeta {
	return &pc.ObjectMeta
}

func (pc *Eye) SetGeneration(generation int64) {
	pc.ObjectMeta.Generation = generation
}

func (pc Eye) GetGeneration() int64 {
	return pc.ObjectMeta.Generation
}

// Registry is an interface for things that know how to store Eye.
// +k8s:deepcopy-gen=false
type EyeRegistry interface {
	ListEyes(ctx request.Context, options *internalversion.ListOptions) (*EyeList, error)
	GetEye(ctx request.Context, id string, options *metav1.GetOptions) (*Eye, error)
	CreateEye(ctx request.Context, id *Eye) (*Eye, error)
	UpdateEye(ctx request.Context, id *Eye) (*Eye, error)
	DeleteEye(ctx request.Context, id string) (bool, error)
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched types will panic.
func NewEyeRegistry(sp builders.StandardStorageProvider) EyeRegistry {
	return &storageEye{sp}
}

// Implement Registry
// storage puts strong typing around storage calls
// +k8s:deepcopy-gen=false
type storageEye struct {
	builders.StandardStorageProvider
}

func (s *storageEye) ListEyes(ctx request.Context, options *internalversion.ListOptions) (*EyeList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	st := s.GetStandardStorage()
	obj, err := st.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*EyeList), err
}

func (s *storageEye) GetEye(ctx request.Context, id string, options *metav1.GetOptions) (*Eye, error) {
	st := s.GetStandardStorage()
	obj, err := st.Get(ctx, id, options)
	if err != nil {
		return nil, err
	}
	return obj.(*Eye), nil
}

func (s *storageEye) CreateEye(ctx request.Context, object *Eye) (*Eye, error) {
	st := s.GetStandardStorage()
	obj, err := st.Create(ctx, object, nil, true)
	if err != nil {
		return nil, err
	}
	return obj.(*Eye), nil
}

func (s *storageEye) UpdateEye(ctx request.Context, object *Eye) (*Eye, error) {
	st := s.GetStandardStorage()
	obj, _, err := st.Update(ctx, object.Name, rest.DefaultUpdatedObjectInfo(object), nil, nil)
	if err != nil {
		return nil, err
	}
	return obj.(*Eye), nil
}

func (s *storageEye) DeleteEye(ctx request.Context, id string) (bool, error) {
	st := s.GetStandardStorage()
	_, sync, err := st.Delete(ctx, id, nil)
	return sync, err
}