package httphandler

import (
	"context"
	"net/http"
	"strings"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/helper"
	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/labstack/echo/v4"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Ambil token dari header
		authHeader := c.Request().Header.Get(echo.HeaderAuthorization)
		splitAuth := strings.Split(authHeader, " ")
		// validate token
		if len(splitAuth) != 2 || splitAuth[0] != "Bearer" {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}
		accessToken := splitAuth[1]

		// Decode token ke claim
		var claim model.CustomClaims
		err := helper.DecodeToken(accessToken, &claim)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		}

		// Kirim ke context
		ctx := context.WithValue(c.Request().Context(), model.BearerAuthKey, claim)
		req := c.Request().WithContext(ctx)
		c.SetRequest(req)

		return next(c)
	}
}
