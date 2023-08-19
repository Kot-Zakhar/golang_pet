package mapper

import (
	"github.com/kot-zakhar/golang_pet/internal/model"
	"github.com/kot-zakhar/golang_pet/internal/viewModel"
)

func MapToViewModel(user *model.User) viewModel.UserViewModel {
	return viewModel.UserViewModel{
		Id:    user.Id,
		Name:  user.Name,
		Login: user.Login,
		Email: user.Email,
	}
}
