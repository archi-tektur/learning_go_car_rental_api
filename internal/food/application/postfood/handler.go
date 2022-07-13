package postfood

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/application"
	"github.com/mtk3d/food-api/internal/food/domain"
)

type Handler struct {
	repository domain.FoodRepository
}

func NewPostFood(repository domain.FoodRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

// GetAllFood
// @Summary      Post food
// @Description  Create new food item
// @Tags         food
// @Accept       json
// @Produce      json
// @Success      201 {string} json "Created"
// @Failure      400 {string} json "Bad request"
// @Failure      404 {string} json "Not found"
// @Failure      500 {string} json "Internal Server Error"
// @Router       /food [post]
func (h *Handler) PostFood(c echo.Context) error {
	foodDTO := new(application.FoodDTO)

	if err := c.Bind(foodDTO); err != nil {
		return c.JSON(http.StatusBadRequest, struct{}{})
	}

	if err := c.Validate(foodDTO); err != nil {
		return err
	}

	h.repository.SaveFood(domain.NewFoodFromDTO(*foodDTO))

	return c.JSON(http.StatusCreated, struct{}{})
}
