package service

import (
	"context"

	"github.com/kot-zakhar/golang_pet/internal/mapper"
	"github.com/kot-zakhar/golang_pet/internal/repository"
	userViewModel "github.com/kot-zakhar/golang_pet/internal/viewModel/user"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (service *UserService) GetAll(context context.Context) (*[]userViewModel.UserViewModel, error) {
	var mappedUsers []userViewModel.UserViewModel
	users, err := service.userRepo.GetAll(context)

	if err != nil {
		return nil, err
	}

	for _, user := range *users {
		mappedUsers = append(mappedUsers, mapper.MapToViewModel(&user))
	}

	return &mappedUsers, nil
}
