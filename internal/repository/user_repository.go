package repository

import (
	"github.com/manish8561/grpc-user-service/internal/model"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	return &model.User{
		ID:    id,
		Name:  "Manish",
		Email: "manish@test.com",
	}, nil
}
