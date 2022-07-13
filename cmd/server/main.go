package server

import (
	"github.com/archi-tektur/car-rental-api/cmd/flags"
	"github.com/archi-tektur/car-rental-api/internal/car/infrastructure/actions"
	"github.com/archi-tektur/car-rental-api/internal/car/infrastructure/storage"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/urfave/cli/v2"
	"log"
)

// application
var app = echo.New()

// global dependencies
var carHandler *actions.CarHandler

func Startup(context *cli.Context) error {
	configure(context)
	configureDependencies()
	registerRoutes()

	log.Fatal(app.Start(":3000"))

	return nil
}

func configure(context *cli.Context) {
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{context.String(flags.CorsAllowOrigin)},
		AllowHeaders: []string{"*"},
	}))
}

func configureDependencies() {
	repository := storage.NewFlatRepository()
	carHandler = actions.NewCarHandler(*repository)
}

func registerRoutes() {
	app.GET("/car", carHandler.ListCars)
	app.GET("/car/:id", carHandler.ShowCar)
	app.POST("/car", carHandler.CreateCar)
}
