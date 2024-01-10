package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmitriysta/messenger/message/internal/pkg/cache"
	"github.com/go-redis/redis/v8"
	"testing"
	"time"

	"github.com/dmitriysta/messenger/message/internal/interfaces/mocks"
	"github.com/dmitriysta/messenger/message/internal/models"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestMessageService_CreateMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	mockCache := new(mocks.RedisClient)
	ctx := context.Background()
	currentTime := time.Now().Truncate(time.Second)

	testMessage := &models.Message{
		UserID:    1,
		ChannelID: 1,
		Content:   "test content",
		CreatedAt: currentTime,
	}

	mockRepo.On("CreateMessage", ctx, mock.AnythingOfType("*models.Message")).Return(nil)
	mockCache.On("Del", ctx, fmt.Sprintf("messages:channel:%d", testMessage.ChannelID)).Return(nil)

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger, mockCache)

	result, err := service.CreateMessage(ctx, testMessage.UserID, testMessage.ChannelID, testMessage.Content)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testMessage.UserID, result.UserID)
	assert.Equal(t, testMessage.ChannelID, result.ChannelID)
	assert.Equal(t, testMessage.Content, result.Content)
	assert.WithinDuration(t, currentTime, result.CreatedAt, time.Second)

	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func TestMessageService_GetMessagesByChannelId(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	mockCache := new(mocks.RedisClient)
	ctx := context.Background()
	currentTime := time.Now().UTC().Truncate(time.Second)

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

	key := fmt.Sprintf("messages:channel:%d", 1)
	jsonData, _ := json.Marshal(testMessages)

	logger, _ := test.NewNullLogger()
	service := NewMessageService(mockRepo, logger, mockCache)
	mockCache.On("Get", ctx, key).Return(redis.NewStringResult(string(jsonData), nil)).Once()
	result, err := service.GetMessagesByChannelId(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, testMessages, result)
	mockCache.AssertCalled(t, "Get", ctx, key)
	mockRepo.AssertNotCalled(t, "GetMessagesByChannelId", ctx, 1)

	mockCache.On("Get", ctx, key).Return(redis.NewStringResult("", redis.Nil)).Once()
	mockRepo.On("GetMessagesByChannelId", ctx, 1).Return(testMessages, nil).Once()
	mockCache.On("Set", ctx, key, jsonData, cache.TimeToLive).Return(redis.NewStatusResult("", nil)).Once()
	result, err = service.GetMessagesByChannelId(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, testMessages, result)
	mockRepo.AssertCalled(t, "GetMessagesByChannelId", ctx, 1)
	mockCache.AssertCalled(t, "Set", ctx, key, jsonData, cache.TimeToLive)

	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
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

	service := NewMessageService(mockRepo, logger, nil)

	result, err := service.GetMessageById(ctx, 1)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, testMessage, result)
	mockRepo.AssertExpectations(t)
}

func TestMessageService_UpdateMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	mockCache := new(mocks.RedisClient)
	ctx := context.Background()

	testMessage := &models.Message{
		UserID:    1,
		ChannelID: 1,
		Content:   "updated content",
	}

	mockRepo.On("UpdateMessage", ctx, testMessage).Return(nil)
	mockCache.On("Del", ctx, fmt.Sprintf("messages:channel:%d", testMessage.ChannelID)).Return(redis.NewIntResult(1, nil))

	logger, _ := test.NewNullLogger()
	service := NewMessageService(mockRepo, logger, mockCache)

	err := service.UpdateMessage(ctx, testMessage)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}

func TestMessageService_DeleteMessage(t *testing.T) {
	mockRepo := new(mocks.MessageRepository)
	mockCache := new(mocks.RedisClient)
	ctx := context.Background()

	mockRepo.On("DeleteMessage", ctx, 1).Return(nil)
	mockCache.On("Del", ctx, fmt.Sprintf("messages:channel:%d", 1)).Return(redis.NewIntResult(1, nil))

	logger, _ := test.NewNullLogger()

	service := NewMessageService(mockRepo, logger, mockCache)

	err := service.DeleteMessage(ctx, 1)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
	mockCache.AssertExpectations(t)
}
