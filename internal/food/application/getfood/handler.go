package getfood

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/application"
	"github.com/mtk3d/food-api/internal/food/domain"
)

type Handler struct {
	repository domain.FoodRepository
}

func NewGetFood(repository domain.FoodRepository) *Handler {
	return &Handler{
		repository: repository,
	}
}

// GetAllFood
// @Summary      Get all food
// @Description  Get all food
// @Tags         food
// @Accept       json
// @Produce      json
// @Success      200 {object} []application.FoodDTO
// @Failure      400 {string} json "Bad request"
// @Failure      404 {string} json "Not found"
// @Failure      500 {string} json "Internal Server Error"
// @Router       /food [get]
func (h *Handler) GetFood(c echo.Context) error {
	allFood := h.repository.GetFood()
	foodDTOs := []application.FoodDTO{}

	for _, food := range allFood {
		foodDTOs = append(foodDTOs, *food.ToDTO())
	}

	return c.JSON(http.StatusOK, foodDTOs)
}
