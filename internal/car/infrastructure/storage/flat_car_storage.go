package storage

import (
	"errors"
	"github.com/archi-tektur/car-rental-api/internal/car/domain"
)

type FlatCarRepository struct {
	cars []domain.Car
}

func NewFlatRepository() *FlatCarRepository {
	r := new(FlatCarRepository)
	return r
}

func (r *FlatCarRepository) Save(car *domain.Car) {
	r.cars = append(r.cars, *car)
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
