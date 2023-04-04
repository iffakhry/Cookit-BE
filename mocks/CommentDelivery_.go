// Code generated by mockery v2.22.1. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"
	mock "github.com/stretchr/testify/mock"
)

// CommentDelivery_ is an autogenerated mock type for the CommentDelivery_ type
type CommentDelivery_ struct {
	mock.Mock
}

// DeleteCommentById provides a mock function with given fields: e
func (_m *CommentDelivery_) DeleteCommentById(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// InsertComment provides a mock function with given fields: e
func (_m *CommentDelivery_) InsertComment(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SelectCommentsByRecipeId provides a mock function with given fields: e
func (_m *CommentDelivery_) SelectCommentsByRecipeId(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateCommentById provides a mock function with given fields: e
func (_m *CommentDelivery_) UpdateCommentById(e echo.Context) error {
	ret := _m.Called(e)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(e)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewCommentDelivery_ interface {
	mock.TestingT
	Cleanup(func())
}

// NewCommentDelivery_ creates a new instance of CommentDelivery_. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCommentDelivery_(t mockConstructorTestingTNewCommentDelivery_) *CommentDelivery_ {
	mock := &CommentDelivery_{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
