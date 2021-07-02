// RAINBOND, Application Management Platform
// Copyright (C) 2014-2021 Goodrain Co., Ltd.

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version. For any non-GPL usage of Rainbond,
// one or multiple Commercial Licenses authorized by Goodrain Co., Ltd.
// must be obtained first.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/goodrain/rainbond/pkg/apis/rainbond/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ComponentDefinitionLister helps list ComponentDefinitions.
// All objects returned here must be treated as read-only.
type ComponentDefinitionLister interface {
	// List lists all ComponentDefinitions in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ComponentDefinition, err error)
	// ComponentDefinitions returns an object that can list and get ComponentDefinitions.
	ComponentDefinitions(namespace string) ComponentDefinitionNamespaceLister
	ComponentDefinitionListerExpansion
}

// componentDefinitionLister implements the ComponentDefinitionLister interface.
type componentDefinitionLister struct {
	indexer cache.Indexer
}

// NewComponentDefinitionLister returns a new ComponentDefinitionLister.
func NewComponentDefinitionLister(indexer cache.Indexer) ComponentDefinitionLister {
	return &componentDefinitionLister{indexer: indexer}
}

// List lists all ComponentDefinitions in the indexer.
func (s *componentDefinitionLister) List(selector labels.Selector) (ret []*v1alpha1.ComponentDefinition, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComponentDefinition))
	})
	return ret, err
}

// ComponentDefinitions returns an object that can list and get ComponentDefinitions.
func (s *componentDefinitionLister) ComponentDefinitions(namespace string) ComponentDefinitionNamespaceLister {
	return componentDefinitionNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ComponentDefinitionNamespaceLister helps list and get ComponentDefinitions.
// All objects returned here must be treated as read-only.
type ComponentDefinitionNamespaceLister interface {
	// List lists all ComponentDefinitions in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.ComponentDefinition, err error)
	// Get retrieves the ComponentDefinition from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.ComponentDefinition, error)
	ComponentDefinitionNamespaceListerExpansion
}

// componentDefinitionNamespaceLister implements the ComponentDefinitionNamespaceLister
// interface.
type componentDefinitionNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ComponentDefinitions in the indexer for a given namespace.
func (s componentDefinitionNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.ComponentDefinition, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.ComponentDefinition))
	})
	return ret, err
}

// Get retrieves the ComponentDefinition from the indexer for a given namespace and name.
func (s componentDefinitionNamespaceLister) Get(name string) (*v1alpha1.ComponentDefinition, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("componentdefinition"), name)
	}
	return obj.(*v1alpha1.ComponentDefinition), nil
}