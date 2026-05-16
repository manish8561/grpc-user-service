package service

import (
	"context"

	"github.com/manish8561/grpc-user-service/internal/model"
	"github.com/manish8561/grpc-user-service/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUser(ctx context.Context, id int64) (*model.User, error) {
	return s.userRepository.FindByID(id)
}
