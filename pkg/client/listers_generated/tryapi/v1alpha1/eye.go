// This file was automatically generated by lister-gen

package v1alpha1

import (
	v1alpha1 "github.com/nightfury1204/test-apiserver/pkg/apis/tryapi/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// EyeLister helps list Eyes.
type EyeLister interface {
	// List lists all Eyes in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Eye, err error)
	// Eyes returns an object that can list and get Eyes.
	Eyes(namespace string) EyeNamespaceLister
	EyeListerExpansion
}

// eyeLister implements the EyeLister interface.
type eyeLister struct {
	indexer cache.Indexer
}

// NewEyeLister returns a new EyeLister.
func NewEyeLister(indexer cache.Indexer) EyeLister {
	return &eyeLister{indexer: indexer}
}

// List lists all Eyes in the indexer.
func (s *eyeLister) List(selector labels.Selector) (ret []*v1alpha1.Eye, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eye))
	})
	return ret, err
}

// Eyes returns an object that can list and get Eyes.
func (s *eyeLister) Eyes(namespace string) EyeNamespaceLister {
	return eyeNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// EyeNamespaceLister helps list and get Eyes.
type EyeNamespaceLister interface {
	// List lists all Eyes in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Eye, err error)
	// Get retrieves the Eye from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Eye, error)
	EyeNamespaceListerExpansion
}

// eyeNamespaceLister implements the EyeNamespaceLister
// interface.
type eyeNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Eyes in the indexer for a given namespace.
func (s eyeNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Eye, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Eye))
	})
	return ret, err
}

// Get retrieves the Eye from the indexer for a given namespace and name.
func (s eyeNamespaceLister) Get(name string) (*v1alpha1.Eye, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("eye"), name)
	}
	return obj.(*v1alpha1.Eye), nil
}