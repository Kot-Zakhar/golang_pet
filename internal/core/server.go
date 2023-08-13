package core

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func SetUpServer(adress string) {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello. This server works fine.")
	})

	e.GET("/users", DI.UserHandler.HandleGetAllUsers)

	e.Logger.Fatal(e.Start(adress))
}
