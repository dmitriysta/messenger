// Code generated by mockery v2.38.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	models "github.com/dmitriysta/messenger/user/internal/models"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *models.User
func (_e *UserRepository_Expecter) CreateUser(ctx interface{}, user interface{}) *UserRepository_CreateUser_Call {
	return &UserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", ctx, user)}
}

func (_c *UserRepository_CreateUser_Call) Run(run func(ctx context.Context, user *models.User)) *UserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserRepository_CreateUser_Call) Return(_a0 error) *UserRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_CreateUser_Call) RunAndReturn(run func(context.Context, *models.User) error) *UserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteUser provides a mock function with given fields: ctx, userId
func (_m *UserRepository) DeleteUser(ctx context.Context, userId int) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_DeleteUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteUser'
type UserRepository_DeleteUser_Call struct {
	*mock.Call
}

// DeleteUser is a helper method to define mock.On call
//   - ctx context.Context
//   - userId int
func (_e *UserRepository_Expecter) DeleteUser(ctx interface{}, userId interface{}) *UserRepository_DeleteUser_Call {
	return &UserRepository_DeleteUser_Call{Call: _e.mock.On("DeleteUser", ctx, userId)}
}

func (_c *UserRepository_DeleteUser_Call) Run(run func(ctx context.Context, userId int)) *UserRepository_DeleteUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *UserRepository_DeleteUser_Call) Return(_a0 error) *UserRepository_DeleteUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_DeleteUser_Call) RunAndReturn(run func(context.Context, int) error) *UserRepository_DeleteUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByEmail provides a mock function with given fields: ctx, email
func (_m *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	ret := _m.Called(ctx, email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*models.User, error)); ok {
		return rf(ctx, email)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *models.User); ok {
		r0 = rf(ctx, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetUserByEmail_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByEmail'
type UserRepository_GetUserByEmail_Call struct {
	*mock.Call
}

// GetUserByEmail is a helper method to define mock.On call
//   - ctx context.Context
//   - email string
func (_e *UserRepository_Expecter) GetUserByEmail(ctx interface{}, email interface{}) *UserRepository_GetUserByEmail_Call {
	return &UserRepository_GetUserByEmail_Call{Call: _e.mock.On("GetUserByEmail", ctx, email)}
}

func (_c *UserRepository_GetUserByEmail_Call) Run(run func(ctx context.Context, email string)) *UserRepository_GetUserByEmail_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *UserRepository_GetUserByEmail_Call) Return(_a0 *models.User, _a1 error) *UserRepository_GetUserByEmail_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetUserByEmail_Call) RunAndReturn(run func(context.Context, string) (*models.User, error)) *UserRepository_GetUserByEmail_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserById provides a mock function with given fields: ctx, userId
func (_m *UserRepository) GetUserById(ctx context.Context, userId int) (*models.User, error) {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for GetUserById")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int) (*models.User, error)); ok {
		return rf(ctx, userId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int) *models.User); ok {
		r0 = rf(ctx, userId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, userId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetUserById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserById'
type UserRepository_GetUserById_Call struct {
	*mock.Call
}

// GetUserById is a helper method to define mock.On call
//   - ctx context.Context
//   - userId int
func (_e *UserRepository_Expecter) GetUserById(ctx interface{}, userId interface{}) *UserRepository_GetUserById_Call {
	return &UserRepository_GetUserById_Call{Call: _e.mock.On("GetUserById", ctx, userId)}
}

func (_c *UserRepository_GetUserById_Call) Run(run func(ctx context.Context, userId int)) *UserRepository_GetUserById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int))
	})
	return _c
}

func (_c *UserRepository_GetUserById_Call) Return(_a0 *models.User, _a1 error) *UserRepository_GetUserById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetUserById_Call) RunAndReturn(run func(context.Context, int) (*models.User, error)) *UserRepository_GetUserById_Call {
	_c.Call.Return(run)
	return _c
}

// UpdateUser provides a mock function with given fields: ctx, user
func (_m *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	ret := _m.Called(ctx, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_UpdateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdateUser'
type UserRepository_UpdateUser_Call struct {
	*mock.Call
}

// UpdateUser is a helper method to define mock.On call
//   - ctx context.Context
//   - user *models.User
func (_e *UserRepository_Expecter) UpdateUser(ctx interface{}, user interface{}) *UserRepository_UpdateUser_Call {
	return &UserRepository_UpdateUser_Call{Call: _e.mock.On("UpdateUser", ctx, user)}
}

func (_c *UserRepository_UpdateUser_Call) Run(run func(ctx context.Context, user *models.User)) *UserRepository_UpdateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.User))
	})
	return _c
}

func (_c *UserRepository_UpdateUser_Call) Return(_a0 error) *UserRepository_UpdateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_UpdateUser_Call) RunAndReturn(run func(context.Context, *models.User) error) *UserRepository_UpdateUser_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
