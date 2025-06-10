package usecase

import (
	"fmt"
	"time"

	"github.com/willians-e-silva/maestro/internal/domain"
)

type UserUsecase struct {
	userRepo domain.UserRepository
}

func NewUserUsecase(ur domain.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: ur}
}

func (uc *UserUsecase) GetUserByID(id string) (*domain.User, error) {
	user, err := uc.userRepo.GetByID(id)
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar usuário: %w", err)
	}
	return user, nil
}

func (uc *UserUsecase) CreateUser(name, email string) (*domain.User, error) {
	existingUser, err := uc.userRepo.GetByEmail(email)
	if err != nil && err != domain.ErrNotFound {
		return nil, fmt.Errorf("falha ao verificar usuário existente: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("usuário com este email já existe")
	}

	user := &domain.User{
		ID:        "some-generated-id",
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("falha ao criar usuário: %w", err)
	}
	return user, nil
}
