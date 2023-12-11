package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dmitriysta/messenger/message/internal/interfaces"
	"github.com/dmitriysta/messenger/message/internal/models"
	"github.com/dmitriysta/messenger/message/internal/pkg/cache"

	"github.com/sirupsen/logrus"
)

const (
	CheMessageChannelPrefix = "messages:channel:"
)

type MessageService struct {
	repo   interfaces.MessageRepository
	logger *logrus.Logger
	cache  interfaces.RedisClient
}

func NewMessageService(repo interfaces.MessageRepository, logger *logrus.Logger, cache interfaces.RedisClient) *MessageService {
	return &MessageService{
		repo:   repo,
		logger: logger,
		cache:  cache,
	}
}

func (s *MessageService) CreateMessage(ctx context.Context, userId, channelId int, content string) (*models.Message, error) {
	message := models.NewMessage(userId, channelId, content)
	if err := s.repo.CreateMessage(ctx, message); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "CreateMessage",
			"error":  err.Error(),
		}).Errorf("failed to create message: %v", err)

		return nil, err
	}

	key := fmt.Sprintf(CheMessageChannelPrefix+"%d", channelId)
	s.cache.Del(ctx, key)

	return message, nil
}

func (s *MessageService) GetMessagesByChannelId(ctx context.Context, channelId int) ([]models.Message, error) {
	key := fmt.Sprintf(CheMessageChannelPrefix+"%d", channelId)

	cachedMessages, err := s.cache.Get(ctx, key).Result()
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessagesByChannelId",
			"error":     err.Error(),
			"channelId": channelId,
		}).Errorf("failed to get messages by channel id: %v", err)

		return nil, err
	} else if err == nil {
		var messages []models.Message
		if err := json.Unmarshal([]byte(cachedMessages), &messages); err != nil {
			s.logger.WithFields(logrus.Fields{
				"module":    "message",
				"func":      "GetMessagesByChannelId",
				"error":     err.Error(),
				"channelId": channelId,
			}).Errorf("failed to unmarshal cached messages: %v", err)

			return nil, err
		}
	}

	messages, err := s.repo.GetMessagesByChannelId(ctx, channelId)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessagesByChannelId",
			"error":     err.Error(),
			"channelId": channelId,
		}).Errorf("failed to get messages by channel id: %v", err)

		return nil, err
	}

	jsonData, err := json.Marshal(messages)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessagesByChannelId",
			"error":     err.Error(),
			"channelId": channelId,
		}).Errorf("failed to marshal messages: %v", err)

		return nil, err
	}

	s.cache.Set(ctx, key, jsonData, cache.TimeToLive)

	return messages, nil
}

func (s *MessageService) GetMessageById(ctx context.Context, messageId int) (*models.Message, error) {
	message, err := s.repo.GetMessageById(ctx, messageId)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessageById",
			"error":     err.Error(),
			"messageId": messageId,
		}).Errorf("failed to get message by id: %v", err)

		return nil, err
	}

	return message, nil
}

func (s *MessageService) UpdateMessage(ctx context.Context, message *models.Message) error {
	if err := s.repo.UpdateMessage(ctx, message); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "UpdateMessage",
			"error":  err.Error(),
		}).Errorf("failed to update message: %v", err)

		return err
	}

	key := fmt.Sprintf(CheMessageChannelPrefix+"%d", message.ChannelID)
	s.cache.Del(ctx, key)

	return nil
}

func (s *MessageService) DeleteMessage(ctx context.Context, messageId int) error {
	if err := s.repo.DeleteMessage(ctx, messageId); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "DeleteMessage",
			"error":  err.Error(),
		}).Errorf("failed to delete message: %v", err)

		return err
	}

	key := fmt.Sprintf(CheMessageChannelPrefix+"%d", messageId)
	s.cache.Del(ctx, key)

	return nil
}
