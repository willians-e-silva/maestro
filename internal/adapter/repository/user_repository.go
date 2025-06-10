package repository

import (
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
