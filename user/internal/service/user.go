package service

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"

	"github.com/dmitriysta/messenger/user/internal/interfaces"
	"github.com/dmitriysta/messenger/user/internal/models"

	"github.com/sirupsen/logrus"
)

type UserService struct {
	repo   interfaces.UserRepository
	logger *logrus.Logger
}

func NewUserService(repo interfaces.UserRepository, logger *logrus.Logger) *UserService {
	return &UserService{
		repo:   repo,
		logger: logger,
	}
}

func (s *UserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {
	user, err := models.NewUser(name, email, password)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":  "user",
			"func":    "CreateUser",
			"subFunc": "NewUser",
			"error":   err.Error(),
		}).Errorf("failed to create user: %v", err)

		return nil, err
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module":  "user",
			"func":    "CreateUser",
			"subFunc": "CreateUser",
			"error":   err.Error(),
		}).Errorf("failed to create user: %v", err)

		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserById(ctx context.Context, userId int) (*models.User, error) {
	user, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "GetUserById",
			"error":  err.Error(),
		}).Errorf("failed to get user by id: %v", err)

		return nil, err
	}

	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, user *models.User) error {
	if err := s.repo.UpdateUser(ctx, user); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "UpdateUser",
			"error":  err.Error(),
		}).Errorf("failed to update user: %v", err)

		return err
	}

	return nil
}

func (s *UserService) DeleteUser(ctx context.Context, userId int) error {
	if err := s.repo.DeleteUser(ctx, userId); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "DeleteUser",
			"error":  err.Error(),
		}).Errorf("failed to delete user: %v", err)

		return err
	}

	return nil
}

func (s *UserService) AuthenticateUser(ctx context.Context, email, password string) (*models.User, string, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "AuthenticateUser",
			"error":  err.Error(),
		}).Errorf("failed to get user by email: %v", err)

		return nil, "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		s.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "AuthenticateUser",
			"error":  err.Error(),
		}).Errorf("failed to authenticate user: %v", err)

		return nil, "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.Id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, "", err
	}

	return user, tokenString, nil
}
