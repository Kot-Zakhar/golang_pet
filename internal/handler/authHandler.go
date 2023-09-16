package handler

import (
	"context"
	"net/http"

	"github.com/kot-zakhar/golang_pet/internal/config"
	"github.com/kot-zakhar/golang_pet/internal/model"
	"github.com/labstack/echo/v4"
)

type IAuthService interface {
	SignIn(context context.Context, login, password, fingerprint, userAgent string) (accessToken string, session model.UserSession, err error)
	SignOut(context context.Context, login string) error
	RefreshTokens(context context.Context, oldRefreshToken string) (accessToken, newRefreshToken string, err error)
}

const defaultApiBaseRoute = "/auth"

type AuthHandler struct {
	authService IAuthService
	apiAuthPath string
	apiDomain   string
}

func NewAuthHandler(config *config.AppConfig, authService IAuthService) AuthHandler {
	return AuthHandler{
		authService: authService,
		apiAuthPath: config.AccessTokenRoute,
		apiDomain:   config.Domain,
	}
}

type LoginInfo struct {
	Login       string `json:"login" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Fingerprint string `json:"fingerprint" validate:"required"`
}

func (handler *AuthHandler) SignIn(c echo.Context) error {
	var loginInfo LoginInfo

	if err := c.Bind(&loginInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if err := c.Validate(&loginInfo); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userAgent := c.Request().UserAgent()

	accessToken, session, err := handler.authService.SignIn(c.Request().Context(), loginInfo.Login, loginInfo.Password, loginInfo.Fingerprint, userAgent)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	cookie := http.Cookie{
		Name:     "refreshToken",
		Value:    session.RefreshToken,
		Domain:   handler.apiDomain,
		Path:     handler.apiAuthPath,
		HttpOnly: true,
		MaxAge:   int(session.ExpiresAt.Sub(session.CreatedAt).Seconds()),
	}

	c.SetCookie(&cookie)

	return c.JSON(http.StatusOK, struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	}{
		accessToken,
		session.RefreshToken,
	})
}

func (handler *AuthHandler) SignOut(c echo.Context) error {
	return echo.ErrNotImplemented
}

func (handler *AuthHandler) RefreshTokens(c echo.Context) error {
	return echo.ErrNotImplemented
}
