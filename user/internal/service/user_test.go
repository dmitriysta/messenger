package service

import (
	"context"
	"github.com/dmitriysta/messenger/user/internal/interfaces/mocks"
	"github.com/dmitriysta/messenger/user/internal/models"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testUser := &models.User{
		Name:      "test username",
		Password:  "test password",
		CreatedAt: currentTime,
	}

	mockRepo.On("CreateUser", ctx, mock.AnythingOfType("*models.User")).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewUserService(mockRepo, logger)

	result, err := service.CreateUser(ctx, testUser.Name, testUser.Email, testUser.Password)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testUser.Name, result.Name)
	assert.Equal(t, testUser.Password, result.Password)
	assert.WithinDuration(t, currentTime, result.CreatedAt, time.Second)
	mockRepo.AssertExpectations(t)
}

func TestUserService_GetUserById(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testUser := &models.User{
		Id:        1,
		Name:      "test username",
		Password:  "test password",
		CreatedAt: currentTime,
	}

	mockRepo.On("GetUserById", ctx, 1).Return(testUser, nil)

	logger, _ := test.NewNullLogger()

	service := NewUserService(mockRepo, logger)

	result, err := service.GetUserById(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testUser.Name, result.Name)
	assert.Equal(t, testUser.Password, result.Password)
	assert.WithinDuration(t, currentTime, result.CreatedAt, time.Second)
	mockRepo.AssertExpectations(t)
}

func TestUserService_UpdateUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testUser := &models.User{
		Name:      "test username",
		Password:  "test password",
		CreatedAt: currentTime,
	}

	mockRepo.On("UpdateUser", ctx, testUser).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewUserService(mockRepo, logger)

	err := service.UpdateUser(ctx, testUser)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := new(mocks.UserRepository)
	ctx := context.Background()

	mockRepo.On("DeleteUser", ctx, 1).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewUserService(mockRepo, logger)

	err := service.DeleteUser(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
