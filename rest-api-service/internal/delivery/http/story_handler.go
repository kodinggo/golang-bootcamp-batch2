package httphandler

import (
	"net/http"

	"github.com/kodinggo/golang-bootcamp-batch2/rest-api-service/internal/model"
	"github.com/labstack/echo/v4"
)

type StoryHandler struct {
	storyUsecase model.StoryUsecase
}

func NewStoryHandler(storyUsecase model.StoryUsecase) *StoryHandler {
	return &StoryHandler{storyUsecase: storyUsecase}
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

	stories, total, err := h.storyUsecase.FindAll(c.Request().Context(), opt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	response := Response{
		Data: stories,
		Metadata: map[string]any{
			"total": total,
		},
	}

	return c.JSON(http.StatusOK, response)
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
