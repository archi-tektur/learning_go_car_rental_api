package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFoodBiting(t *testing.T) {
	food := &Food{
		id:     1,
		name:   "Pizza",
		price:  "5$",
		weight: 10,
	}

	bite := NewBite(1)

	food.Bite(bite)

	foodDto := food.ToDTO()

	assert.Equal(t, int64(9), foodDto.Weight, "Pizza should be bitten")
}
