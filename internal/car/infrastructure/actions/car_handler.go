package actions

import "github.com/archi-tektur/car-rental-api/internal/car/infrastructure/storage"

type CarHandler struct {
	repository storage.FlatCarRepository
}

func NewCarHandler(repository storage.FlatCarRepository) *CarHandler {
	return &CarHandler{
		repository: repository,
	}
}
