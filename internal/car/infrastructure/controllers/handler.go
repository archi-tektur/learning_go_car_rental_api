package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/car/domain"
	"net/http"
)

type Handler struct{}

func NewListCarsHandler() *Handler {
	return &Handler{}
}

func (h *Handler) ListCars(context echo.Context) error {
	cars := []domain.Car{
		{"Mercedes GLC", "WW 4545L"},
		{"Seat Leon", "WW 4646L"},
	}

	return context.JSON(http.StatusOK, cars)
}
