package models

import (
	"errors"
	"time"
)

const (
	ErrInvalidUserID    = "invalid user id"
	ErrInvalidChannelID = "invalid channel id"
	ErrInvalidContent   = "invalid content"
)

type Message struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    int       `json:"userId" gorm:"index:idx_user_id"`
	ChannelID int       `json:"channelId" gorm:"index:idx_channel_id"`
	Content   string    `json:"content" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"-" gorm:"autoDeleteTime"`
}

func (m *Message) Validate() error {
	if m.UserID == 0 {
		return errors.New(ErrInvalidUserID)
	}

	if m.ChannelID == 0 {
		return errors.New(ErrInvalidChannelID)
	}

	if m.Content == "" {
		return errors.New(ErrInvalidContent)
	}

	return nil
}

func NewMessage(userId, channelId int, content string) *Message {
	return &Message{
		UserID:    userId,
		ChannelID: channelId,
		Content:   content,
		CreatedAt: time.Now(),
	}
}
