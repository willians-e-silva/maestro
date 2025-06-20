package userusecase

import (
	"context"
	"fmt"
	"time"

	errors "maestro/internal/domain/errors"
	user "maestro/internal/domain/user"

	pb "maestro/internal/infra/grpc/user"
)

type UserUsecase struct {
	userRepo user.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserUsecase(ur user.UserRepository) *UserUsecase {
	return &UserUsecase{userRepo: ur}
}

func (uc *UserUsecase) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.UserResponse, error) {
	user, err := uc.userRepo.GetByID(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar usuário por ID: %w", err)
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uc *UserUsecase) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.UserResponse, error) {
	existingUser, err := uc.userRepo.GetByEmail(req.GetEmail())
	if err != nil && err != errors.ErrNotFound {
		return nil, fmt.Errorf("falha ao verificar usuário existente: %w", err)
	}
	if existingUser != nil {
		return nil, fmt.Errorf("usuário com este email já existe")
	}

	user := &user.User{
		ID:        "some-generated-id",
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, fmt.Errorf("falha ao criar usuário: %w", err)
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uc *UserUsecase) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	err := uc.userRepo.Delete(ctx, req.GetId())
	if err != nil {
		return nil, fmt.Errorf("falha ao deletar usuário: %w", err)
	}

	return &pb.DeleteUserResponse{
		Success: true,
	}, nil
}

func (uc *UserUsecase) GetUserByEmail(ctx context.Context, req *pb.GetUserByEmailRequest) (*pb.UserResponse, error) {
	user, err := uc.userRepo.GetByEmail(req.GetEmail())
	if err != nil {
		return nil, fmt.Errorf("falha ao buscar usuário por email: %w", err)
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (uc *UserUsecase) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UserResponse, error) {
	user := &user.User{
		ID:        req.GetId(),
		Name:      req.GetName(),
		Email:     req.GetEmail(),
		UpdatedAt: time.Now(),
	}

	if err := uc.userRepo.Update(user); err != nil {
		return nil, fmt.Errorf("falha ao atualizar usuário: %w", err)
	}

	return &pb.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
