// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import account "github.com/Optum/dce/pkg/account"
import mock "github.com/stretchr/testify/mock"

// Writer is an autogenerated mock type for the Writer type
type Writer struct {
	mock.Mock
}

// Write provides a mock function with given fields: i, lastModifiedOn
func (_m *Writer) Write(i *account.Account, lastModifiedOn *int64) error {
	ret := _m.Called(i, lastModifiedOn)

	var r0 error
	if rf, ok := ret.Get(0).(func(*account.Account, *int64) error); ok {
		r0 = rf(i, lastModifiedOn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}