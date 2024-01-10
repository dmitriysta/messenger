package repository

import (
	"context"
	"database/sql"

	"github.com/dmitriysta/messenger/message/internal/models"

	"github.com/sirupsen/logrus"
)

type MessageRepository struct {
	db     *sql.DB
	logger *logrus.Logger
}

func NewMessageRepository(db *sql.DB, logger *logrus.Logger) *MessageRepository {
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

	query := `INSERT INTO message (user_id, channel_id, content, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	err := r.db.QueryRowContext(ctx, query, message.UserID, message.ChannelID, message.Content, message.CreatedAt).Scan(&message.Id)
	if err != nil {
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
	query := `SELECT * FROM message WHERE channel_id = $1`
	rows, err := r.db.QueryContext(ctx, query, channelId)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessagesByChannelId",
			"error":     err.Error(),
			"channelId": channelId,
		}).Errorf("failed to get messages by channel id: %v", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.Id, &message.UserID, &message.ChannelID, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.DeletedAt)
		if err != nil {
			r.logger.WithFields(logrus.Fields{
				"module":    "message",
				"func":      "GetMessagesByChannelId",
				"error":     err.Error(),
				"channelId": channelId,
			}).Errorf("failed to scan message: %v", err)
			return nil, err
		}

		messages = append(messages, message)
	}

	if err := rows.Err(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessagesByChannelId",
			"error":     err.Error(),
			"channelId": channelId,
		}).Errorf("failed to get messages by channel id: %v", err)
		return nil, err
	}

	return messages, nil
}

func (r *MessageRepository) GetMessageById(ctx context.Context, messageId int) (*models.Message, error) {
	var message models.Message
	query := `SELECT * FROM message WHERE id = $1`
	err := r.db.QueryRowContext(ctx, query, messageId).Scan(&message.Id, &message.UserID, &message.ChannelID, &message.Content, &message.CreatedAt, &message.UpdatedAt, &message.DeletedAt)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module":    "message",
			"func":      "GetMessageById",
			"error":     err.Error(),
			"messageId": messageId,
		}).Errorf("failed to get message by id: %v", err)

		return nil, err
	}

	return &message, nil
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

	query := `UPDATE message SET user_id = $1, channel_id = $2, content = $3, updated_at = $4 WHERE id = $5`
	_, err := r.db.ExecContext(ctx, query, message.UserID, message.ChannelID, message.Content, message.UpdatedAt, message.Id)
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "message",
			"func":   "UpdateMessage",
			"error":  err.Error(),
		}).Errorf("failed to update message: %v", err)

		return err
	}

	return nil
}

func (r *MessageRepository) DeleteMessage(ctx context.Context, messageId int) error {
	query := `DELETE FROM message WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, messageId)
	if err != nil {
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
