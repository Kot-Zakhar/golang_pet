package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/kot-zakhar/golang_pet/internal/service"
	userViewModel "github.com/kot-zakhar/golang_pet/internal/viewModel/user"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) HandleGetAllUsers(c echo.Context) error {
	users, err := handler.userService.GetAll(c.Request().Context())

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	usersResponse := userViewModel.GetAllResponseViewModel{
		Users: *users,
	}

	return c.JSON(http.StatusOK, usersResponse)
}
