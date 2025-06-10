package graphql

import (
	"context"
	"fmt"

	"github.com/willians-e-silva/maestro/internal/domain"
	"github.com/willians-e-silva/maestro/internal/usecase"
)

type Resolver struct {
	UserUsecase *usecase.UserUsecase // Injete seus casos de uso aqui
}

// NewResolver cria um novo Resolver
func NewResolver(userUC *usecase.UserUsecase) *Resolver {
	return &Resolver{
		UserUsecase: userUC,
	}
}

func (r *mutationResolver) CreateUser(ctx context.Context, input NewUserInput) (*User, error) {
	user, err := r.Resolver.UserUsecase.CreateUser(input.Name, input.Email)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar usuário: %w", err)
	}
	// Mapear domain.User para graphql.User (se necessário, o gqlgen pode fazer isso automaticamente para campos compatíveis)
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (r *queryResolver) User(ctx context.Context, id string) (*User, error) {
	user, err := r.Resolver.UserUsecase.GetUserByID(id)
	if err != nil {
		if err == domain.ErrNotFound { // Se você definiu este erro
			return nil, fmt.Errorf("usuário não encontrado")
		}
		return nil, fmt.Errorf("erro ao buscar usuário: %w", err)
	}
	// Mapear domain.User para graphql.User
	return &User{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
