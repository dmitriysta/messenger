//go:generate mockery

package interfaces

import (
	"context"
	"message/internal/models"
)

type MessageService interface {
	CreateMessage(ctx context.Context, userId, channelId int, content string) (*models.Message, error)
	GetMessagesByChannelId(ctx context.Context, channelId int) ([]models.Message, error)
	GetMessageById(ctx context.Context, messageId int) (*models.Message, error)
	UpdateMessage(ctx context.Context, message *models.Message) error
	DeleteMessage(ctx context.Context, messageId int) error
}
