package postbite

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/application"
	"github.com/mtk3d/food-api/internal/food/domain"
)

type Handler struct {
	repository domain.FoodRepository
}

func NewPostBite(repository domain.FoodRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

// GetAllFood
// @Summary      Post food bite
// @Description  Bite a selected food
// @Tags         food
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Food ID"
// @Success      200 {string} json "OK"
// @Failure      400 {string} json "Bad request"
// @Failure      404 {string} json "Not found"
// @Failure      500 {string} json "Internal Server Error"
// @Router       /food/{id}/bite [post]
func (h *Handler) PostBite(c echo.Context) error {
	biteDTO := new(application.BiteDTO)

	if err := c.Bind(biteDTO); err != nil {
		return c.JSON(http.StatusBadRequest, struct{}{})
	}

	if err := c.Validate(biteDTO); err != nil {
		return err
	}

	bite := domain.NewBite(biteDTO.Weight)

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct{}{})
	}

	food, err := h.repository.FindFood(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, struct{}{})
	}

	food.Bite(bite)

	h.repository.SaveFood(food)

	return c.JSON(http.StatusOK, struct{}{})
}
