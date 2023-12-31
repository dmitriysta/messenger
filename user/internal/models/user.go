package models

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"time"
)

const (
	ErrInvalidName     = "invalid name"
	ErrInvalidEmail    = "invalid email"
	ErrInvalidPassword = "invalid password"
)

type User struct {
	Id        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255)"`
	Email     string    `json:"email" gorm:"type:varchar(255);unique"`
	Password  string    `json:"password" gorm:"type:varchar(255)"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
	DeletedAt time.Time `json:"-" gorm:"autoDeleteTime"`
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New(ErrInvalidName)
	}

	if u.Email == "" {
		return errors.New(ErrInvalidEmail)
	}

	if u.Password == "" {
		return errors.New(ErrInvalidPassword)
	}

	return nil
}

func NewUser(name, email, password string) (*User, error) {
	hashedPassword, err := hashPassword(password)
	if err != nil {
		return nil, err
	}

	return &User{
		Name:      name,
		Email:     email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
