// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	runtime "github.com/go-openapi/runtime"
	mock "github.com/stretchr/testify/mock"

	usergroup "github.com/mittwald/goharbor-client/v5/apiv2/internal/api/client/usergroup"
)

// MockUsergroupClientService is an autogenerated mock type for the ClientService type
type MockUsergroupClientService struct {
	mock.Mock
}

// CreateUserGroup provides a mock function with given fields: params, authInfo
func (_m *MockUsergroupClientService) CreateUserGroup(params *usergroup.CreateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*usergroup.CreateUserGroupCreated, error) {
	ret := _m.Called(params, authInfo)

	var r0 *usergroup.CreateUserGroupCreated
	if rf, ok := ret.Get(0).(func(*usergroup.CreateUserGroupParams, runtime.ClientAuthInfoWriter) *usergroup.CreateUserGroupCreated); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usergroup.CreateUserGroupCreated)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usergroup.CreateUserGroupParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteUserGroup provides a mock function with given fields: params, authInfo
func (_m *MockUsergroupClientService) DeleteUserGroup(params *usergroup.DeleteUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*usergroup.DeleteUserGroupOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *usergroup.DeleteUserGroupOK
	if rf, ok := ret.Get(0).(func(*usergroup.DeleteUserGroupParams, runtime.ClientAuthInfoWriter) *usergroup.DeleteUserGroupOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usergroup.DeleteUserGroupOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usergroup.DeleteUserGroupParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserGroup provides a mock function with given fields: params, authInfo
func (_m *MockUsergroupClientService) GetUserGroup(params *usergroup.GetUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*usergroup.GetUserGroupOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *usergroup.GetUserGroupOK
	if rf, ok := ret.Get(0).(func(*usergroup.GetUserGroupParams, runtime.ClientAuthInfoWriter) *usergroup.GetUserGroupOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usergroup.GetUserGroupOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usergroup.GetUserGroupParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListUserGroups provides a mock function with given fields: params, authInfo
func (_m *MockUsergroupClientService) ListUserGroups(params *usergroup.ListUserGroupsParams, authInfo runtime.ClientAuthInfoWriter) (*usergroup.ListUserGroupsOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *usergroup.ListUserGroupsOK
	if rf, ok := ret.Get(0).(func(*usergroup.ListUserGroupsParams, runtime.ClientAuthInfoWriter) *usergroup.ListUserGroupsOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usergroup.ListUserGroupsOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usergroup.ListUserGroupsParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// SetTransport provides a mock function with given fields: transport
func (_m *MockUsergroupClientService) SetTransport(transport runtime.ClientTransport) {
	_m.Called(transport)
}

// UpdateUserGroup provides a mock function with given fields: params, authInfo
func (_m *MockUsergroupClientService) UpdateUserGroup(params *usergroup.UpdateUserGroupParams, authInfo runtime.ClientAuthInfoWriter) (*usergroup.UpdateUserGroupOK, error) {
	ret := _m.Called(params, authInfo)

	var r0 *usergroup.UpdateUserGroupOK
	if rf, ok := ret.Get(0).(func(*usergroup.UpdateUserGroupParams, runtime.ClientAuthInfoWriter) *usergroup.UpdateUserGroupOK); ok {
		r0 = rf(params, authInfo)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*usergroup.UpdateUserGroupOK)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*usergroup.UpdateUserGroupParams, runtime.ClientAuthInfoWriter) error); ok {
		r1 = rf(params, authInfo)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
