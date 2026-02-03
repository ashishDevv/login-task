package repository

import (
	"context"

	"github.com/aashisDevv/login-api/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
