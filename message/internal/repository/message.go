package repository

import (
	"context"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"message/internal/models"
)

type MessageRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewMessageRepository(db *gorm.DB, logger *logrus.Logger) *MessageRepository {
	return &MessageRepository{
		db:     db,
		logger: logger,
	}
}

func (r *MessageRepository) CreateMessage(ctx context.Context, message *models.Message) error {
	if err := message.Validate(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "CreateMessage",
			"error":  err.Error(),
		}).Errorf("failed to validate message: %v", err)

		return err
	}

	if err := r.db.WithContext(ctx).Create(message).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "CreateMessage",
			"error":  err.Error(),
		}).Errorf("failed to create message: %v", err)

		return err
	}

	return nil
}

func (r *MessageRepository) GetMessagesByChannelId(ctx context.Context, channelId int) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.WithContext(ctx).Where("channel_id = ?", channelId).Find(&messages).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "GetMessagesByChannelId",
			"error":  err.Error(),
		}).Errorf("failed to get messages by channel id: %v", err)
	}

	return messages, err
}

func (r *MessageRepository) GetMessageById(ctx context.Context, messageId int) (*models.Message, error) {
	var message models.Message
	err := r.db.WithContext(ctx).Where("id = ?", messageId).First(&message).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessageById",
			"error":     err.Error(),
			"messageId": messageId,
		}).Errorf("failed to get message by id: %v", err)
		return nil, err
	}

	return &message, err
}

func (r *MessageRepository) UpdateMessage(ctx context.Context, message *models.Message) error {
	if err := message.Validate(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "UpdateMessage",
			"error":  err.Error(),
		}).Errorf("failed to validate message: %v", err)

		return err
	}

	if err := r.db.WithContext(ctx).Save(message).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "UpdateMessage",
			"error":     err.Error(),
			"messageId": message.Id,
		}).Errorf("failed to update message: %v", err)

		return err
	}

	return nil
}

func (r *MessageRepository) DeleteMessage(ctx context.Context, messageId int) error {
	if err := r.db.WithContext(ctx).Model(&models.Message{}).Where("id = ?", messageId).Delete(&models.Message{}).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "DeleteMessage",
			"error":     err.Error(),
			"messageId": messageId,
		}).Errorf("failed to delete message: %v", err)

		return err
	}

	return nil
}
