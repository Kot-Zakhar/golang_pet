package service

import (
	"context"
	"fmt"

	"github.com/kot-zakhar/golang_pet/internal/handler"
	"github.com/kot-zakhar/golang_pet/internal/model"
)

type IUserRepository interface {
	GetAll(context.Context) ([]model.User, error)
	GetById(context.Context, uint64) (model.User, error)
	Insert(context.Context, model.User) error
	Update(context.Context, uint64, model.User) error
	Delete(context.Context, uint64) error
}

type UserService struct {
	userRepo IUserRepository
}

func NewUserService(userRepo IUserRepository) handler.IUserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func (service *UserService) GetAll(context context.Context) ([]model.User, error) {
	users, err := service.userRepo.GetAll(context)

	if err != nil {
		return nil, fmt.Errorf("UserService:GetAll:service.userRepo.GetAll - %w", err)
	}

	return users, nil
}

func (service *UserService) GetById(context context.Context, id uint64) (model.User, error) {
	user, err := service.userRepo.GetById(context, id)
	if err != nil {
		return user, fmt.Errorf("UserService:GetAll:service.userRepo.GetById - %w", err)
	}

	return user, nil
}

func (service *UserService) RegisterUser(context context.Context, registeredUser model.User) error {
	// existingUser, err := service.userRepo.GetByLogin(context, createUserModel.Login)
	// if err != nil {
	// 	return fmt.Errorf("UserService:RegisterUser:service.userRepo.GetByLogin - %w", err)
	// }

	// if existingUser != nil {
	// 	return fmt.Errorf("User was already registered.")
	// }

	err := service.userRepo.Insert(context, registeredUser)

	if err != nil {
		return fmt.Errorf("UserService:RegisterUser:service.userRepo.Insert - %w", err)
	}

	return nil
}

func (service *UserService) UpdateUser(context context.Context, id uint64, updateUserModel model.User) error {
	existingUser, err := service.userRepo.GetById(context, id)
	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.GetById - %w", err)
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
	_, err := service.userRepo.GetById(context, id)
	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.GetById - %w", err)
	}

	err = service.userRepo.Delete(context, id)

	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.Delete - %w", err)
	}

	return nil
}
