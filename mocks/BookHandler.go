// Code generated by mockery v2.40.1. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// BookHandler is an autogenerated mock type for the BookHandler type
type BookHandler struct {
	mock.Mock
}

// CreateBook provides a mock function with given fields: c
func (_m *BookHandler) CreateBook(c *gin.Context) {
	_m.Called(c)
}

// DeleteBook provides a mock function with given fields: c
func (_m *BookHandler) DeleteBook(c *gin.Context) {
	_m.Called(c)
}

// GetBook provides a mock function with given fields: c
func (_m *BookHandler) GetBook(c *gin.Context) {
	_m.Called(c)
}

// GetBooks provides a mock function with given fields: c
func (_m *BookHandler) GetBooks(c *gin.Context) {
	_m.Called(c)
}

// UpdateBook provides a mock function with given fields: c
func (_m *BookHandler) UpdateBook(c *gin.Context) {
	_m.Called(c)
}

// NewBookHandler creates a new instance of BookHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookHandler {
	mock := &BookHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}