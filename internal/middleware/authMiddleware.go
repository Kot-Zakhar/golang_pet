package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/labstack/echo/v4"
)

type IJwtValidatorService interface {
	ValidateAndGetUserId(tokenString string) (userId string, err error)
}

func NewAuthMiddleware(config *config.AppConfig, jwtValidator IJwtValidatorService) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := strings.Split(c.Request().Header.Get("Authorization"), "Bearer ")
			if len(authHeader) != 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Malformed token")
			}

			token := authHeader[1]
			userId, err := jwtValidator.ValidateAndGetUserId(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, fmt.Errorf("Unauthorized - %w", err))
			}

			c.Set("userId", userId)

			return next(c)
		}
	}
}
