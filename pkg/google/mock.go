// Code generated by mockery v2.42.3. DO NOT EDIT.

package google

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	multipart "mime/multipart"
)

// Mock is an autogenerated mock type for the Google type
type Mock struct {
	mock.Mock
}

// Delete provides a mock function with given fields: ctx, productName
func (_m *Mock) Delete(ctx context.Context, productName string) error {
	ret := _m.Called(ctx, productName)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, productName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upload provides a mock function with given fields: ctx, product, file
func (_m *Mock) Upload(ctx context.Context, product string, file multipart.File) error {
	ret := _m.Called(ctx, product, file)

	if len(ret) == 0 {
		panic("no return value specified for Upload")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, multipart.File) error); ok {
		r0 = rf(ctx, product, file)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMock creates a new instance of Mock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMock(t interface {
	mock.TestingT
	Cleanup(func())
}) *Mock {
	mock := &Mock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
