package main

import (
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

	e.GET("/users", func(c echo.Context) error {
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
	})
	e.POST("/users", func(c echo.Context) error {
		var body requestBody
		if err := c.Bind(&body); err != nil {
			log.Println(err)
			return echo.NewHTTPError(http.StatusBadRequest, "error bind body")
		}

		fmt.Println("body Name", body.Name)

		return c.JSON(http.StatusCreated, nil)
	})
	e.POST("/users/:id", func(c echo.Context) error {
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
	})

	e.Start(":1323")
}
