package service

import (
	"context"
	"github.com/sirupsen/logrus"
	"message/internal/interfaces"
	"message/internal/models"
)

type MessageService struct {
	repo   interfaces.MessageRepository
	logger *logrus.Logger
}

func NewMessageService(repo interfaces.MessageRepository, logger *logrus.Logger) *MessageService {
	return &MessageService{
		repo:   repo,
		logger: logger,
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

	return message, nil
}

func (s *MessageService) GetMessagesByChannelId(ctx context.Context, channelId int) ([]models.Message, error) {
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

	return nil
}
