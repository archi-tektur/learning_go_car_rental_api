package actions

import (
	"github.com/archi-tektur/car-rental-api/internal/car/infrastructure/storage"
	"github.com/archi-tektur/car-rental-api/internal/shared/httperrors"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (this *CarHandler) CreateCar(context echo.Context) error {
	createCarData := new(storage.CreateCarData)

	if err := context.Bind(createCarData); err != nil {
		message := httperrors.CarCreationFailedDueUserInput()
		return context.JSON(message.Code, message)
	}

	id := this.repository.Save(createCarData)

	car, _ := this.repository.FindCarById(id)

	return context.JSON(http.StatusOK, car)
}
