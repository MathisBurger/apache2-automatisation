package main

import (
	"github.com/MathisBurger/apache2-automatisation/config"
	"github.com/MathisBurger/apache2-automatisation/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	cfg, err := config.ParseConfig()
	if err != nil {
		panic(err)
	}

	app := fiber.New()
	app.Use(cors.New())

	app.Use(logger.New())

	app.Get("/configureWordpress", controller.ConfigureWordpressController)

	app.Listen(":" + cfg.ApplicationPort)
}
