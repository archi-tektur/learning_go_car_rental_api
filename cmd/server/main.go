// @title           food-api
// @version         1.0
// @description     Example food api to learn golang
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8080
// @BasePath  /
package server

import (
    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
    "github.com/mtk3d/food-api/cmd/flags"
    "github.com/mtk3d/food-api/internal/car/infrastructure/controllers"
    "github.com/urfave/cli/v2"
    "log"
)

func Command(c *cli.Context) error {
    app := echo.New()

    app.Use(middleware.Logger())
    app.Use(middleware.Recover())

    app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
        AllowOrigins: []string{c.String(flags.CorsAllowOrigin)},
        AllowHeaders: []string{"*"},
    }))

    app.GET("/api", func(context *echo.Context) error {
        return c.SendString("Hello, World!")
    })

    app.GET("/car", controllers.NewListCarsHandler().ListCars)

    //app.Get("/food", controllers.NewListCarsHandler().ListCars)
    //e.POST("/food", postfood.NewPostFood(repository).PostFood)
    //e.POST("/food/:id/bite", postbite.NewPostBite(repository).PostBite)

    log.Fatal(app.Listen(":3000"))

    return nil
}
