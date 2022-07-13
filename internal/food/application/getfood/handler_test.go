package getfood

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/mtk3d/food-api/internal/food/domain"
	"github.com/mtk3d/food-api/internal/food/infrastructure/db"
	"github.com/mtk3d/food-api/internal/validator"
	"github.com/stretchr/testify/assert"
)

var (
	foodJSON = `[{"id":1,"name":"Pizza","price":"5$","weight":10}]` + "\n"
)

func TestGetFood(t *testing.T) {
	// Given
	e := echo.New()
	e.Validator = validator.NewValidator()
	repository := db.NewInMemoryFoodRepository()

	repository.SaveFood(domain.NewFood(1, "Pizza", "5$", 10))

	// When
	getFoodReq := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(""))
	getFoodReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	getFoodRec := httptest.NewRecorder()
	getFoodContext := e.NewContext(getFoodReq, getFoodRec)
	getFoodContext.SetPath("/food")

	getFoodHandler := NewGetFood(repository)
	getFoodHandler.GetFood(getFoodContext)

	// Then
	assert.Equal(t, foodJSON, getFoodRec.Body.String())
}
