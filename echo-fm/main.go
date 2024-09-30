package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type requestParam struct {
	ID string `param:"id"`
}

type requestQueryStr struct {
	Search string `query:"search"`
	SortBy string `query:"sort_by"`
}

type requestBody struct {
	Name string `json:"name"`
}

func main() {
	e := echo.New()

	// Middleware
	trackMiddleware := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			log.Println("before entering handler")

			ctx := context.WithValue(c.Request().Context(), "user_id", 1)
			req := c.Request().WithContext(ctx)
			c.SetRequest(req)

			err := next(c)

			log.Println("after entering handler")

			return err
		}
	}

	e.GET("/users", findUsers, trackMiddleware)
	e.POST("/users", createUser)
	e.PUT("/users/:id", updateUser)

	e.Start(":1323")
}

func findUsers(c echo.Context) error {
	log.Println("your are in handler")

	userID := c.Request().Context().Value("user_id")
	fmt.Println("user_id from context", userID)

	var qs requestQueryStr
	if err := c.Bind(&qs); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "error bind qs")
	}

	fmt.Println("qs Search", qs.Search)
	fmt.Println("qs SortBy", qs.SortBy)

	data := map[string]any{
		"id":    1,
		"title": "lorem ipsum",
	}
	return c.JSON(http.StatusOK, data)
}

func createUser(c echo.Context) error {
	var body requestBody
	if err := c.Bind(&body); err != nil {
		log.Println(err)
		return echo.NewHTTPError(http.StatusBadRequest, "error bind body")
	}

	fmt.Println("body Name", body.Name)

	return c.JSON(http.StatusCreated, nil)
}

func updateUser(c echo.Context) error {
	var param requestParam
	if err := c.Bind(&param); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "error bind param")
	}
	fmt.Println("param ID", param.ID)

	data := map[string]any{
		"id":    1,
		"title": "lorem ipsum",
	}
	return c.JSON(http.StatusOK, data)
}
