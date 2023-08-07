package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"golang_pet/app/handler"
)

func Host() {
	e := echo.New()

	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello. This server works fine.")
	})

	e.GET("/events", handler.GetEvents)

	e.Logger.Fatal(e.Start(":1323"))
}
