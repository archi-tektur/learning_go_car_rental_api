package domain

type FoodRepository interface {
	SaveFood(*Food)
	GetFood() []Food
	FindFood(id int64) (*Food, error)
}
