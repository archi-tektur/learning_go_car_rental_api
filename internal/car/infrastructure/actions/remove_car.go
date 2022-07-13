package actions

import (
	"github.com/archi-tektur/car-rental-api/internal/shared/httperrors"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (this *CarHandler) RemoveCar(context echo.Context) error {

	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		message := httperrors.ProvidedIdIsnNotAnInteger()
		return context.JSON(message.Code, message)
	}

	car, _ := this.repository.FindCarById(id)

	if car == nil {
		message := httperrors.CarNotFound()
		return context.JSON(message.Code, message)
	}

	this.repository.RemoveCar(id)

	return context.JSON(http.StatusNoContent, nil)
}
