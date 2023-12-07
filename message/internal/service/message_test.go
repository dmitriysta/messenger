package service

import (
	"context"
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"message/internal/interfaces/mocks"
	"message/internal/models"
	"testing"
	"time"
)

func TestMessageService_CreateMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testMessage := &models.Message{
		UserID:    1,
		ChannelID: 1,
		Content:   "test content",
		CreatedAt: currentTime,
	}

	mockRepo.On("CreateMessage", ctx, mock.AnythingOfType("*models.Message")).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger)

	result, err := service.CreateMessage(ctx, testMessage.UserID, testMessage.ChannelID, testMessage.Content)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testMessage.UserID, result.UserID)
	assert.Equal(t, testMessage.ChannelID, result.ChannelID)
	assert.Equal(t, testMessage.Content, result.Content)
	assert.WithinDuration(t, currentTime, result.CreatedAt, time.Second)
	mockRepo.AssertExpectations(t)
}

func TestMessageService_GetMessagesByChannelId(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testMessages := []models.Message{
		{
			UserID:    1,
			ChannelID: 1,
			Content:   "test content 1",
			CreatedAt: currentTime,
		},
		{
			UserID:    2,
			ChannelID: 1,
			Content:   "test content 2",
			CreatedAt: currentTime,
		},
	}

	mockRepo.On("GetMessagesByChannelId", ctx, 1).Return(testMessages, nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger)

	result, err := service.GetMessagesByChannelId(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testMessages, result)
	mockRepo.AssertExpectations(t)
}

func TestMessageService_GetMessageById(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testMessage := &models.Message{
		UserID:    1,
		ChannelID: 1,
		Content:   "test content",
		CreatedAt: currentTime,
	}

	mockRepo.On("GetMessageById", ctx, 1).Return(testMessage, nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger)

	result, err := service.GetMessageById(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testMessage, result)
	mockRepo.AssertExpectations(t)
}

func TestMessageService_UpdateMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testMessage := &models.Message{
		UserID:    1,
		ChannelID: 1,
		Content:   "test content",
		CreatedAt: currentTime,
	}

	mockRepo.On("UpdateMessage", ctx, testMessage).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger)

	err := service.UpdateMessage(ctx, testMessage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestMessageService_DeleteMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	ctx := context.Background()

	mockRepo.On("DeleteMessage", ctx, 1).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger)

	err := service.DeleteMessage(ctx, 1)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}
