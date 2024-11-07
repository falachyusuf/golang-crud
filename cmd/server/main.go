package main

import (
	"github.com/falachyusuf_/golang-crud/cmd/router"
	"github.com/falachyusuf_/golang-crud/internal/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"github.com/gofiber/fiber/v2/middleware/recover"
)


func main() {
	app := fiber.New()
	app.Use(recover.New())
	app.Use(cache.New())

	// Add application route
	router.Routers(app)

	port := config.DoGetServerPort()
	if err := app.Listen(":" + port); err != nil {
		panic(err)
	}
}
