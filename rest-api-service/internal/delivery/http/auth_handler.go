package httphandler

import (
	"net/http"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	authUsecase model.AuthUsecase
}

func NewAuthHandler(authUsecase model.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

func (h *AuthHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/auth")

	g.POST("/login", h.login)
	g.POST("/register", h.register)
}

func (h *AuthHandler) login(c echo.Context) error {
	var body model.Login
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := h.authUsecase.Login(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}

func (h *AuthHandler) register(c echo.Context) error {
	var body model.Register
	if err := c.Bind(&body); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	accessToken, err := h.authUsecase.Register(c.Request().Context(), body)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Response{
		AccessToken: accessToken,
	})
}
