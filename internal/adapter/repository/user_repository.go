package repository

import (
	"context"

	errors "github.com/willians-e-silva/maestro/internal/domain/errors"
	user "github.com/willians-e-silva/maestro/internal/domain/user"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id string) (*user.User, error) {
	var user user.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Create(user *user.User) error {
	return r.db.Create(user).Error
}

func (r *PostgresUserRepository) Delete(ctx context.Context, userID string) error {
	result := r.db.WithContext(ctx).Delete(&user.User{}, "id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *PostgresUserRepository) GetByEmail(email string) (*user.User, error) {
	var user user.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Update(user *user.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
