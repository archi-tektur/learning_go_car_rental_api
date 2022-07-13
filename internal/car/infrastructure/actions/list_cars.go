package actions

import (
	"github.com/archi-tektur/car-rental-api/internal/car/domain"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (this *CarHandler) ListCars(context echo.Context) error {
	cars := this.repository.FindAllCars()

	if len(cars) == 0 {
		cars = make([]domain.Car, 0)
	}

	return context.JSON(http.StatusOK, cars)
}
