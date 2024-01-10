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
	Id        int       `json:"id"`
	UserID    int       `json:"userId"`
	ChannelID int       `json:"channelId"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"-"`
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
