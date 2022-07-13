package domain

import "github.com/mtk3d/food-api/internal/food/application"

type Food struct {
	id     int64
	name   string
	price  string
	weight int64
}

func NewFood(id int64, name string, price string, weight int64) *Food {
	return &Food{
		id:     id,
		name:   name,
		price:  price,
		weight: weight,
	}
}

func NewFoodFromDTO(dto application.FoodDTO) *Food {
	return NewFood(dto.Id, dto.Name, dto.Price, dto.Weight)
}

func (f *Food) Bite(bite Bite) {
	f.weight -= bite.GetWeight()
}

func (f *Food) ToDTO() *application.FoodDTO {
	return &application.FoodDTO{
		Id:     f.id,
		Name:   f.name,
		Price:  f.price,
		Weight: f.weight,
	}
}
