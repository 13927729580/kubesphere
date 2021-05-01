// Copyright 2018 The prometheus-operator Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ThanosRulerLister helps list ThanosRulers.
type ThanosRulerLister interface {
	// List lists all ThanosRulers in the indexer.
	List(selector labels.Selector) (ret []*v1.ThanosRuler, err error)
	// ThanosRulers returns an object that can list and get ThanosRulers.
	ThanosRulers(namespace string) ThanosRulerNamespaceLister
	ThanosRulerListerExpansion
}

// thanosRulerLister implements the ThanosRulerLister interface.
type thanosRulerLister struct {
	indexer cache.Indexer
}

// NewThanosRulerLister returns a new ThanosRulerLister.
func NewThanosRulerLister(indexer cache.Indexer) ThanosRulerLister {
	return &thanosRulerLister{indexer: indexer}
}

// List lists all ThanosRulers in the indexer.
func (s *thanosRulerLister) List(selector labels.Selector) (ret []*v1.ThanosRuler, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ThanosRuler))
	})
	return ret, err
}

// ThanosRulers returns an object that can list and get ThanosRulers.
func (s *thanosRulerLister) ThanosRulers(namespace string) ThanosRulerNamespaceLister {
	return thanosRulerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ThanosRulerNamespaceLister helps list and get ThanosRulers.
type ThanosRulerNamespaceLister interface {
	// List lists all ThanosRulers in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.ThanosRuler, err error)
	// Get retrieves the ThanosRuler from the indexer for a given namespace and name.
	Get(name string) (*v1.ThanosRuler, error)
	ThanosRulerNamespaceListerExpansion
}

// thanosRulerNamespaceLister implements the ThanosRulerNamespaceLister
// interface.
type thanosRulerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all ThanosRulers in the indexer for a given namespace.
func (s thanosRulerNamespaceLister) List(selector labels.Selector) (ret []*v1.ThanosRuler, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.ThanosRuler))
	})
	return ret, err
}

// Get retrieves the ThanosRuler from the indexer for a given namespace and name.
func (s thanosRulerNamespaceLister) Get(name string) (*v1.ThanosRuler, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("thanosruler"), name)
	}
	return obj.(*v1.ThanosRuler), nil
}
