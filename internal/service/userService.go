package service

import (
	"context"
	"fmt"

	"github.com/kot-zakhar/golang_pet/internal/model"
)

type IUserRepository interface {
	GetAll(context.Context) ([]model.User, error)
	GetById(context.Context, uint64) (model.User, error)
	Insert(context.Context, model.User) (model.User, error)
	Update(context.Context, uint64, model.User) error
	Delete(context.Context, uint64) error
}

type IPasswordHasherService interface {
	HashPassword(password string) (hash, salt []byte)
}

type UserService struct {
	userRepo      IUserRepository
	cryptoService IPasswordHasherService
}

func NewUserService(userRepo IUserRepository, cryptoService IPasswordHasherService) UserService {
	return UserService{userRepo, cryptoService}
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

type UserRegistrationInfo struct {
	Name     string
	Login    string
	Email    string
	Password string
}

func (service *UserService) RegisterUser(context context.Context, registrationInfo UserRegistrationInfo) error {
	hash, salt := service.cryptoService.HashPassword(registrationInfo.Password)

	user := model.User{
		Name:         registrationInfo.Name,
		Login:        registrationInfo.Login,
		PasswordHash: hash,
		Salt:         salt,
		Email:        registrationInfo.Email,
	}

	user, err := service.userRepo.Insert(context, user)

	if err != nil {
		return fmt.Errorf("UserService:RegisterUser:service.userRepo.Insert - %w", err)
	}

	return nil
}

func (service *UserService) UpdateUser(context context.Context, id uint64, updateUserModel UserRegistrationInfo) error {
	existingUser, err := service.userRepo.GetById(context, id)
	if err != nil {
		return fmt.Errorf("UserService:UpdateUser:service.userRepo.GetById - %w", err)
	}

	// TODO: move password change to a separate endpoint/service_method
	hash, salt := service.cryptoService.HashPassword(updateUserModel.Password)

	existingUser.Login = updateUserModel.Login
	existingUser.Email = updateUserModel.Email
	existingUser.Name = updateUserModel.Name
	existingUser.PasswordHash = hash
	existingUser.Salt = salt

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
