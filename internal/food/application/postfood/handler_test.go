package postfood

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/application"
	"github.com/mtk3d/food-api/internal/food/infrastructure/db"
	"github.com/mtk3d/food-api/internal/validator"
	"github.com/stretchr/testify/assert"
)

var postFoodJSON = `{"id":1,"name":"Pizza","price":"5$","weight":10}`

func TestCreateFood(t *testing.T) {
	// Given
	e := echo.New()
	e.Validator = validator.NewValidator()
	repository := db.NewInMemoryFoodRepository()

	// When
	postFoodReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(postFoodJSON))
	postFoodReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	postFoodRec := httptest.NewRecorder()
	postFoodContext := e.NewContext(postFoodReq, postFoodRec)
	postFoodContext.SetPath("/food")

	postFoodHandler := NewPostFood(repository)
	postFoodHandler.PostFood(postFoodContext)

	// Then
	food, _ := repository.FindFood(1)

	assert.NotNil(t, food)

	assert.Equal(t, &application.FoodDTO{
		Id:     1,
		Name:   "Pizza",
		Price:  "5$",
		Weight: 10,
	}, food.ToDTO())
}
