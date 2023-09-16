package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/kot-zakhar/golang_pet/internal/model"
	"github.com/kot-zakhar/golang_pet/internal/service"
)

type IUserService interface {
	GetAll(context.Context) ([]model.User, error)
	GetById(context.Context, uint64) (model.User, error)
	RegisterUser(context.Context, service.UserRegistrationInfo) error
	UpdateUser(context.Context, uint64, service.UserRegistrationInfo) error
	DeleteUser(context.Context, uint64) error
}

type UserHandler struct {
	userService IUserService
}

func NewUserHandler(service IUserService) UserHandler {
	return UserHandler{
		userService: service,
	}
}

type CreateOrUpdateUserViewModel struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserViewModel struct {
	Id    uint64 `json:"id" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Login string `json:"login" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

func (handler *UserHandler) GetAllUsers(c echo.Context) error {
	users, err := handler.userService.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userViewModels := make([]UserViewModel, 0)

	for _, u := range users {
		userViewModels = append(userViewModels, modelToViewModel(u))
	}

	return c.JSON(http.StatusOK, userViewModels)
}

func (handler *UserHandler) GetUserById(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user, err := handler.userService.GetById(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	userViewModel := modelToViewModel(user)

	return c.JSON(http.StatusOK, userViewModel)
}

func (handler *UserHandler) CreateUser(c echo.Context) error {
	// validate user record
	var userModel CreateOrUpdateUserViewModel

	if err := c.Bind(&userModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err := c.Validate(&userModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err := handler.userService.RegisterUser(c.Request().Context(), viewModelToDto(userModel))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (handler *UserHandler) UpdateUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var userModel CreateOrUpdateUserViewModel

	if err = c.Bind(&userModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if err = c.Validate(&userModel); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = handler.userService.UpdateUser(c.Request().Context(), id, viewModelToDto(userModel))
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (handler *UserHandler) DeleteUser(c echo.Context) error {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = handler.userService.DeleteUser(c.Request().Context(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func viewModelToDto(viewModel CreateOrUpdateUserViewModel) service.UserRegistrationInfo {
	return service.UserRegistrationInfo{
		Name:     viewModel.Name,
		Login:    viewModel.Login,
		Email:    viewModel.Email,
		Password: viewModel.Password,
	}
}

func modelToViewModel(user model.User) UserViewModel {
	return UserViewModel{
		Id:    user.Id,
		Name:  user.Name,
		Login: user.Login,
		Email: user.Email,
	}
}
