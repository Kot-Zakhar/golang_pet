package core

import (
	"net/http"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		// Optionally, you could return the error to give each route more control over the status code
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetUpServer(adress string) {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello. This server works fine.")
	})

	userHandler := DI.UserHandler
	userBaseRoute := "/api/users"

	e.GET(userBaseRoute, userHandler.GetAllUsers)
	e.POST(userBaseRoute, userHandler.CreateUser)
	e.GET(userBaseRoute+"/:id", userHandler.GetUserById)
	e.PUT(userBaseRoute+"/:id", userHandler.UpdateUser)
	e.DELETE(userBaseRoute+"/:id", userHandler.DeleteUser)

	authHandler := DI.AuthHandler
	authBaseRoute := "/api/auth"

	e.POST(authBaseRoute+"/sign-in", authHandler.SignIn)
	e.POST(authBaseRoute+"/sign-out", authHandler.SignOut, DI.AuthMiddleware)
	e.POST(authBaseRoute+"/refresh-tokens", authHandler.RefreshTokens) // Q - do we need AuthMiddleware here? Maybe some kind of RefreshTokens(allowExpired:true)

	e.Logger.Fatal(e.Start(adress))
}
