package application

type FoodDTO struct {
	Id     int64  `json:"id" form:"id" validate:"required,gte=0"`
	Name   string `json:"name" form:"name" validate:"required"`
	Price  string `json:"price" form:"price" validate:"required,endswith=$"`
	Weight int64  `json:"weight" form:"weight" validate:"required,gte=0"`
}

type BiteDTO struct {
	Weight int64 `json:"weight" form:"weight" validate:"required,gte=0"`
}
