//go:generate mockery

package interfaces

import (
	"context"

	"github.com/dmitriysta/messenger/message/internal/models"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, message *models.Message) error
	GetMessagesByChannelId(ctx context.Context, channelId int) ([]models.Message, error)
	GetMessageById(ctx context.Context, messageId int) (*models.Message, error)
	UpdateMessage(ctx context.Context, message *models.Message) error
	DeleteMessage(ctx context.Context, messageId int) error
}
