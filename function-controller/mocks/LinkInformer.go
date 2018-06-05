// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import cache "k8s.io/client-go/tools/cache"
import mock "github.com/stretchr/testify/mock"
import projectriffv1alpha1 "github.com/projectriff/riff/kubernetes-crds/pkg/client/listers/projectriff.io/v1alpha1"

// LinkInformer is an autogenerated mock type for the LinkInformer type
type LinkInformer struct {
	mock.Mock
}

// Informer provides a mock function with given fields:
func (_m *LinkInformer) Informer() cache.SharedIndexInformer {
	ret := _m.Called()

	var r0 cache.SharedIndexInformer
	if rf, ok := ret.Get(0).(func() cache.SharedIndexInformer); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cache.SharedIndexInformer)
		}
	}

	return r0
}

// Lister provides a mock function with given fields:
func (_m *LinkInformer) Lister() projectriffv1alpha1.LinkLister {
	ret := _m.Called()

	var r0 projectriffv1alpha1.LinkLister
	if rf, ok := ret.Get(0).(func() projectriffv1alpha1.LinkLister); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(projectriffv1alpha1.LinkLister)
		}
	}

	return r0
}
