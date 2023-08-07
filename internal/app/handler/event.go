package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"golang_pet/app/model"
	"golang_pet/app/service"
)

func GetEvents(c echo.Context) error {
	// Some validation

	var events []model.Event = service.GetEvents()

	return c.JSON(http.StatusOK, events)
}
