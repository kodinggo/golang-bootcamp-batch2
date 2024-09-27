package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/users", func(c echo.Context) error {
		data := map[string]any{
			"id":    1,
			"title": "lorem ipsum",
		}
		return c.JSON(http.StatusOK, data)
	})
	e.POST("/users", func(c echo.Context) error {
		return c.JSON(http.StatusCreated, nil)
	})

	e.Start(":1323")
}
