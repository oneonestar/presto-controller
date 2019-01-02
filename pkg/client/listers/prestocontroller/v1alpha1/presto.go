/*
Copyright 2019 oneonestar.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/oneonestar/presto-controller/pkg/apis/prestocontroller/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// PrestoLister helps list Prestos.
type PrestoLister interface {
	// List lists all Prestos in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Presto, err error)
	// Prestos returns an object that can list and get Prestos.
	Prestos(namespace string) PrestoNamespaceLister
	PrestoListerExpansion
}

// prestoLister implements the PrestoLister interface.
type prestoLister struct {
	indexer cache.Indexer
}

// NewPrestoLister returns a new PrestoLister.
func NewPrestoLister(indexer cache.Indexer) PrestoLister {
	return &prestoLister{indexer: indexer}
}

// List lists all Prestos in the indexer.
func (s *prestoLister) List(selector labels.Selector) (ret []*v1alpha1.Presto, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Presto))
	})
	return ret, err
}

// Prestos returns an object that can list and get Prestos.
func (s *prestoLister) Prestos(namespace string) PrestoNamespaceLister {
	return prestoNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// PrestoNamespaceLister helps list and get Prestos.
type PrestoNamespaceLister interface {
	// List lists all Prestos in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Presto, err error)
	// Get retrieves the Presto from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Presto, error)
	PrestoNamespaceListerExpansion
}

// prestoNamespaceLister implements the PrestoNamespaceLister
// interface.
type prestoNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Prestos in the indexer for a given namespace.
func (s prestoNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Presto, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Presto))
	})
	return ret, err
}

// Get retrieves the Presto from the indexer for a given namespace and name.
func (s prestoNamespaceLister) Get(name string) (*v1alpha1.Presto, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("presto"), name)
	}
	return obj.(*v1alpha1.Presto), nil
}
