package storage

import (
	"errors"
	domain "github.com/archi-tektur/car-rental-api/internal/car/domain"
)

type CreateCarData struct {
	Model              string `json:"model"`
	RegistrationPlates string `json:"registration_plates"`
}

type FlatCarRepository struct {
	cars []domain.Car
}

func NewFlatRepository() *FlatCarRepository {
	r := new(FlatCarRepository)
	return r
}

func (r *FlatCarRepository) Save(request *CreateCarData) int {
	carId := r.findLastIndex() + 1

	car := domain.Car{
		Id:                 carId,
		Model:              request.Model,
		RegistrationPlates: request.RegistrationPlates,
	}

	r.cars = append(r.cars, car)

	return carId
}

func (r *FlatCarRepository) RemoveFood(id int) {
	found, _ := r.GetCarById(id)

	if found == nil {
		return
	}

	var sliceWithoutRemovedElement []domain.Car
	for _, car := range r.cars {
		if car.Id != found.Id {
			sliceWithoutRemovedElement = append(sliceWithoutRemovedElement, car)
		}
	}
	r.cars = sliceWithoutRemovedElement
}

func (r *FlatCarRepository) GetAllCars() []domain.Car {
	return r.cars
}

func (r *FlatCarRepository) findLastIndex() int {
	if len(r.cars) == 0 {
		return 0
	}

	var min int = 0
	for _, car := range r.cars {
		if car.Id > min {
			min = car.Id
		}
	}

	return min
}

func (r *FlatCarRepository) GetCarById(id int) (*domain.Car, error) {
	var found *domain.Car

	for _, car := range r.cars {
		if car.Id == id {
			found = &car
		}
	}

	if found == nil {
		return nil, errors.New("car not found")
	}

	return found, nil
}
