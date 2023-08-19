package service

import (
	"context"
	"fmt"

	"github.com/kot-zakhar/golang_pet/internal/mapper"
	"github.com/kot-zakhar/golang_pet/internal/model"
	"github.com/kot-zakhar/golang_pet/internal/repository"
	"github.com/kot-zakhar/golang_pet/internal/viewModel"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (service *UserService) GetAll(context context.Context) (*[]viewModel.UserViewModel, error) {
	mappedUsers := make([]viewModel.UserViewModel, 0)
	users, err := service.userRepo.GetAll(context)

	if err != nil {
		return nil, fmt.Errorf("UserService:GetAll:service.userRepo.GetAll - %w", err)
	}

	for _, user := range *users {
		mappedUsers = append(mappedUsers, mapper.MapToViewModel(&user))
	}

	return &mappedUsers, nil
}

func (service *UserService) GetById(context context.Context, id uint64) (*viewModel.UserViewModel, error) {
	user, err := service.userRepo.GetById(context, id)
	if err != nil {
		return nil, fmt.Errorf("UserService:GetAll:service.userRepo.GetById - %w", err)
	}

	mappedUser := mapper.MapToViewModel(user)

	return &mappedUser, nil
}

func (service *UserService) RegisterUser(context context.Context, createUserModel *viewModel.CreateOrUpdateUserViewModel) error {
	// existingUser, err := service.userRepo.GetByLogin(context, createUserModel.Login)
	// if err != nil {
	// 	return fmt.Errorf("UserService:RegisterUser:service.userRepo.GetByLogin - %w", err)
	// }

	// if existingUser != nil {
	// 	return fmt.Errorf("User was already registered.")
	// }

	registeredUser := model.User{
		Name:     createUserModel.Name,
		Login:    createUserModel.Login,
		Password: createUserModel.Password,
		Email:    createUserModel.Email,
	}

	_, err := service.userRepo.Insert(context, &registeredUser)

	if err != nil {
		return fmt.Errorf("UserService:RegisterUser:service.userRepo.Insert - %w", err)
	}

	return nil
}

func (service *UserService) UpdateUser(context context.Context, id uint64, updateUserModel *viewModel.CreateOrUpdateUserViewModel) error {
	existingUser, err := service.userRepo.GetById(context, id)
	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.GetById - %w", err)
	}

	if existingUser == nil {
		return fmt.Errorf("User was not found.")
	}

	existingUser.Login = updateUserModel.Login
	existingUser.Email = updateUserModel.Email
	existingUser.Name = updateUserModel.Name
	existingUser.Password = updateUserModel.Password

	err = service.userRepo.Update(context, id, existingUser)

	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.Update - %w", err)
	}

	return nil
}

func (service *UserService) DeleteUser(context context.Context, id uint64) error {
	existingUser, err := service.userRepo.GetById(context, id)
	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.GetById - %w", err)
	}

	if existingUser == nil {
		return fmt.Errorf("User was not found.")
	}

	err = service.userRepo.Delete(context, id)

	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.Delete - %w", err)
	}

	return nil
}
