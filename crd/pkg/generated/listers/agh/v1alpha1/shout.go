/*
Copyright The Kubernetes Authors.

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
	v1alpha1 "github.com/bskiba/agh-talk/crd/pkg/apis/agh/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ShoutLister helps list Shouts.
type ShoutLister interface {
	// List lists all Shouts in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Shout, err error)
	// Shouts returns an object that can list and get Shouts.
	Shouts(namespace string) ShoutNamespaceLister
	ShoutListerExpansion
}

// shoutLister implements the ShoutLister interface.
type shoutLister struct {
	indexer cache.Indexer
}

// NewShoutLister returns a new ShoutLister.
func NewShoutLister(indexer cache.Indexer) ShoutLister {
	return &shoutLister{indexer: indexer}
}

// List lists all Shouts in the indexer.
func (s *shoutLister) List(selector labels.Selector) (ret []*v1alpha1.Shout, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Shout))
	})
	return ret, err
}

// Shouts returns an object that can list and get Shouts.
func (s *shoutLister) Shouts(namespace string) ShoutNamespaceLister {
	return shoutNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// ShoutNamespaceLister helps list and get Shouts.
type ShoutNamespaceLister interface {
	// List lists all Shouts in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Shout, err error)
	// Get retrieves the Shout from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Shout, error)
	ShoutNamespaceListerExpansion
}

// shoutNamespaceLister implements the ShoutNamespaceLister
// interface.
type shoutNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Shouts in the indexer for a given namespace.
func (s shoutNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Shout, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Shout))
	})
	return ret, err
}

// Get retrieves the Shout from the indexer for a given namespace and name.
func (s shoutNamespaceLister) Get(name string) (*v1alpha1.Shout, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("shout"), name)
	}
	return obj.(*v1alpha1.Shout), nil
}