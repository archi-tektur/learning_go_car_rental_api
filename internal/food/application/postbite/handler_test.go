package postbite

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/application"
	"github.com/mtk3d/food-api/internal/food/domain"
	"github.com/mtk3d/food-api/internal/food/infrastructure/db"
	"github.com/mtk3d/food-api/internal/validator"
	"github.com/stretchr/testify/assert"
)

var (
	biteJSON = `{"weight":1}`
)

func TestBitingFood(t *testing.T) {
	// Given
	e := echo.New()
	e.Validator = validator.NewValidator()
	repository := db.NewInMemoryFoodRepository()
	repository.SaveFood(domain.NewFood(1, "Pizza", "5$", 10))

	// When
	postBiteReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(biteJSON))
	postBiteReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	postBiteRec := httptest.NewRecorder()
	postBiteContext := e.NewContext(postBiteReq, postBiteRec)
	postBiteContext.SetPath("food/:id/bite")
	postBiteContext.SetParamNames("id")
	postBiteContext.SetParamValues("1")

	postBiteHandler := NewPostBite(repository)
	postBiteHandler.PostBite(postBiteContext)

	// Then
	food, _ := repository.FindFood(1)

	assert.Equal(t, &application.FoodDTO{
		Id:     1,
		Name:   "Pizza",
		Price:  "5$",
		Weight: 9,
	}, food.ToDTO())
}
