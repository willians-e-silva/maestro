package repository

import (
	"context"

	"github.com/willians-e-silva/maestro/internal/domain"

	"gorm.io/gorm"
)

type PostgresUserRepository struct {
	db *gorm.DB // ou *sql.DB
}

func NewPostgresUserRepository(db *gorm.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) GetByID(id string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "id = ?", id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound // Mapear erro GORM para erro de domínio
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Create(user *domain.User) error {
	return r.db.Create(user).Error
}

// Implementação do método Delete com context.Context
func (r *PostgresUserRepository) Delete(ctx context.Context, userID string) error {
	result := r.db.WithContext(ctx).Delete(&domain.User{}, "id = ?", userID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Implementação do método GetByEmail
func (r *PostgresUserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, "email = ?", email).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, domain.ErrNotFound // Mapear erro GORM para erro de domínio
		}
		return nil, err
	}
	return &user, nil
}

func (r *PostgresUserRepository) Update(user *domain.User) error {
	result := r.db.Save(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
