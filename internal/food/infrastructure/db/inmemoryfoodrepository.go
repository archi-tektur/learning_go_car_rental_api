package db

import (
	"errors"

	"github.com/mtk3d/food-api/internal/food/domain"
)

type InMemoryFoodRepository struct {
	food []domain.Food
}

func NewInMemoryFoodRepository() *InMemoryFoodRepository {
	r := new(InMemoryFoodRepository)
	return r
}

func (r *InMemoryFoodRepository) SaveFood(food *domain.Food) {
	r.RemoveFood(food.ToDTO().Id)
	r.food = append(r.food, *food)
}

func (r *InMemoryFoodRepository) RemoveFood(id int64) {
	found, _ := r.FindFood(id)

	if found != nil {
		var newSlice []domain.Food
		for _, food := range r.food {
			if food.ToDTO().Id != found.ToDTO().Id {
				newSlice = append(newSlice, food)
			}
		}
		r.food = newSlice
	}
}

func (r *InMemoryFoodRepository) GetFood() []domain.Food {
	return r.food
}

func (r *InMemoryFoodRepository) FindFood(id int64) (*domain.Food, error) {
	var found *domain.Food

	for _, food := range r.food {
		if food.ToDTO().Id == id {
			found = &food
		}
	}

	if found == nil {
		return nil, errors.New("Food not found")
	}

	return found, nil
}
