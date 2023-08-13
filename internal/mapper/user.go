package mapper

import (
	"github.com/kot-zakhar/golang_pet/internal/model"
	userViewModel "github.com/kot-zakhar/golang_pet/internal/viewModel/user"
)

func MapToViewModel(user *model.User) userViewModel.UserViewModel {
	return userViewModel.UserViewModel{
		Id:    user.Id,
		Name:  user.Name,
		Login: user.Login,
	}
}
