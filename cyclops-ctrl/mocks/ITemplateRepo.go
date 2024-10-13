// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	models "github.com/cyclops-ui/cyclops/cyclops-ctrl/internal/models"
	ristretto "github.com/dgraph-io/ristretto"
	mock "github.com/stretchr/testify/mock"

	v1alpha1 "github.com/cyclops-ui/cyclops/cyclops-ctrl/api/v1alpha1"
)

// ITemplateRepo is an autogenerated mock type for the ITemplateRepo type
type ITemplateRepo struct {
	mock.Mock
}

type ITemplateRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *ITemplateRepo) EXPECT() *ITemplateRepo_Expecter {
	return &ITemplateRepo_Expecter{mock: &_m.Mock}
}

// GetTemplate provides a mock function with given fields: repo, path, version, resolvedVersion, source
func (_m *ITemplateRepo) GetTemplate(repo string, path string, version string, resolvedVersion string, source v1alpha1.TemplateSourceType) (*models.Template, error) {
	ret := _m.Called(repo, path, version, resolvedVersion, source)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplate")
	}

	var r0 *models.Template
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, string, v1alpha1.TemplateSourceType) (*models.Template, error)); ok {
		return rf(repo, path, version, resolvedVersion, source)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, string, v1alpha1.TemplateSourceType) *models.Template); ok {
		r0 = rf(repo, path, version, resolvedVersion, source)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.Template)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, string, v1alpha1.TemplateSourceType) error); ok {
		r1 = rf(repo, path, version, resolvedVersion, source)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ITemplateRepo_GetTemplate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplate'
type ITemplateRepo_GetTemplate_Call struct {
	*mock.Call
}

// GetTemplate is a helper method to define mock.On call
//   - repo string
//   - path string
//   - version string
//   - resolvedVersion string
//   - source v1alpha1.TemplateSourceType
func (_e *ITemplateRepo_Expecter) GetTemplate(repo interface{}, path interface{}, version interface{}, resolvedVersion interface{}, source interface{}) *ITemplateRepo_GetTemplate_Call {
	return &ITemplateRepo_GetTemplate_Call{Call: _e.mock.On("GetTemplate", repo, path, version, resolvedVersion, source)}
}

func (_c *ITemplateRepo_GetTemplate_Call) Run(run func(repo string, path string, version string, resolvedVersion string, source v1alpha1.TemplateSourceType)) *ITemplateRepo_GetTemplate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(string), args[4].(v1alpha1.TemplateSourceType))
	})
	return _c
}

func (_c *ITemplateRepo_GetTemplate_Call) Return(_a0 *models.Template, _a1 error) *ITemplateRepo_GetTemplate_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ITemplateRepo_GetTemplate_Call) RunAndReturn(run func(string, string, string, string, v1alpha1.TemplateSourceType) (*models.Template, error)) *ITemplateRepo_GetTemplate_Call {
	_c.Call.Return(run)
	return _c
}

// GetTemplateInitialValues provides a mock function with given fields: repo, path, version, source
func (_m *ITemplateRepo) GetTemplateInitialValues(repo string, path string, version string, source v1alpha1.TemplateSourceType) (map[string]interface{}, error) {
	ret := _m.Called(repo, path, version, source)

	if len(ret) == 0 {
		panic("no return value specified for GetTemplateInitialValues")
	}

	var r0 map[string]interface{}
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string, string, v1alpha1.TemplateSourceType) (map[string]interface{}, error)); ok {
		return rf(repo, path, version, source)
	}
	if rf, ok := ret.Get(0).(func(string, string, string, v1alpha1.TemplateSourceType) map[string]interface{}); ok {
		r0 = rf(repo, path, version, source)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]interface{})
		}
	}

	if rf, ok := ret.Get(1).(func(string, string, string, v1alpha1.TemplateSourceType) error); ok {
		r1 = rf(repo, path, version, source)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ITemplateRepo_GetTemplateInitialValues_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetTemplateInitialValues'
type ITemplateRepo_GetTemplateInitialValues_Call struct {
	*mock.Call
}

// GetTemplateInitialValues is a helper method to define mock.On call
//   - repo string
//   - path string
//   - version string
//   - source v1alpha1.TemplateSourceType
func (_e *ITemplateRepo_Expecter) GetTemplateInitialValues(repo interface{}, path interface{}, version interface{}, source interface{}) *ITemplateRepo_GetTemplateInitialValues_Call {
	return &ITemplateRepo_GetTemplateInitialValues_Call{Call: _e.mock.On("GetTemplateInitialValues", repo, path, version, source)}
}

func (_c *ITemplateRepo_GetTemplateInitialValues_Call) Run(run func(repo string, path string, version string, source v1alpha1.TemplateSourceType)) *ITemplateRepo_GetTemplateInitialValues_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string), args[2].(string), args[3].(v1alpha1.TemplateSourceType))
	})
	return _c
}

func (_c *ITemplateRepo_GetTemplateInitialValues_Call) Return(_a0 map[string]interface{}, _a1 error) *ITemplateRepo_GetTemplateInitialValues_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *ITemplateRepo_GetTemplateInitialValues_Call) RunAndReturn(run func(string, string, string, v1alpha1.TemplateSourceType) (map[string]interface{}, error)) *ITemplateRepo_GetTemplateInitialValues_Call {
	_c.Call.Return(run)
	return _c
}

// ReturnCache provides a mock function with given fields:
func (_m *ITemplateRepo) ReturnCache() *ristretto.Cache {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for ReturnCache")
	}

	var r0 *ristretto.Cache
	if rf, ok := ret.Get(0).(func() *ristretto.Cache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ristretto.Cache)
		}
	}

	return r0
}

// ITemplateRepo_ReturnCache_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReturnCache'
type ITemplateRepo_ReturnCache_Call struct {
	*mock.Call
}

// ReturnCache is a helper method to define mock.On call
func (_e *ITemplateRepo_Expecter) ReturnCache() *ITemplateRepo_ReturnCache_Call {
	return &ITemplateRepo_ReturnCache_Call{Call: _e.mock.On("ReturnCache")}
}

func (_c *ITemplateRepo_ReturnCache_Call) Run(run func()) *ITemplateRepo_ReturnCache_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *ITemplateRepo_ReturnCache_Call) Return(_a0 *ristretto.Cache) *ITemplateRepo_ReturnCache_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *ITemplateRepo_ReturnCache_Call) RunAndReturn(run func() *ristretto.Cache) *ITemplateRepo_ReturnCache_Call {
	_c.Call.Return(run)
	return _c
}

// NewITemplateRepo creates a new instance of ITemplateRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewITemplateRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *ITemplateRepo {
	mock := &ITemplateRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
