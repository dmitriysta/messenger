package repository

import (
	"context"

	"github.com/dmitriysta/messenger/user/internal/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(db *gorm.DB, logger *logrus.Logger) *UserRepository {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "CreateUser",
			"error":  err.Error(),
		}).Errorf("failed to validate user: %v", err)

		return err
	}

	if err := r.db.WithContext(ctx).Create(user).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "CreateUser",
			"error":  err.Error(),
		}).Errorf("failed to create user: %v", err)

		return err
	}

	return nil
}

func (r *UserRepository) GetUserById(ctx context.Context, userId int) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", userId).First(&user).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "GetUserById",
			"error":  err.Error(),
		}).Errorf("failed to get user by id: %v", err)

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "GetUserByEmail",
			"error":  err.Error(),
		}).Errorf("failed to get user by email: %v", err)

		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, user *models.User) error {
	if err := user.Validate(); err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "UpdateUser",
			"error":  err.Error(),
		}).Errorf("failed to validate user: %v", err)

		return err
	}

	if err := r.db.WithContext(ctx).Save(user).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "UpdateUser",
			"error":  err.Error(),
		}).Errorf("failed to update user: %v", err)

		return err
	}

	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userId int) error {
	if err := r.db.WithContext(ctx).Delete(&models.User{}, userId).Error; err != nil {
		r.logger.WithFields(logrus.Fields{
			"module": "user",
			"func":   "DeleteUser",
			"error":  err.Error(),
		}).Errorf("failed to delete user: %v", err)

		return err
	}

	return nil
}
