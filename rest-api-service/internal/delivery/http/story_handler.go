package httphandler

import (
	"net/http"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/labstack/echo/v4"
)

type StoryHandler struct {
	StoryUsecase model.StoryUsecase
}

func NewStoryHandler() *StoryHandler {
	return &StoryHandler{}
}

func (h *StoryHandler) RegisterRoutes(e *echo.Echo) {
	g := e.Group("/stories")

	g.GET("", h.findAll)
	g.GET("/:id", h.findByID)
	g.POST("", h.create)
	g.PUT("/:id", h.update)
	g.DELETE("/:id", h.delete)
}

func (h *StoryHandler) findAll(c echo.Context) error {
	var opt model.StoryOpt
	if err := c.Bind(&opt); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	// TODO
	// stories, total, err := h.StoryUsecase.FindAll(c.Request().Context(), opt)

	return c.JSON(http.StatusOK, nil)
}

func (h *StoryHandler) findByID(c echo.Context) error {
	// TODO
	return c.JSON(http.StatusOK, nil)
}

func (h *StoryHandler) create(c echo.Context) error {
	// TODO
	return c.JSON(http.StatusCreated, nil)
}

func (h *StoryHandler) update(c echo.Context) error {
	// TODO
	return c.JSON(http.StatusOK, nil)
}

func (h *StoryHandler) delete(c echo.Context) error {
	// TODO
	return c.NoContent(http.StatusNoContent)
}
