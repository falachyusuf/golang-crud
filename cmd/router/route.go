package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func Routers(app *fiber.App) {
	app.Use(cors.New())
	app.Get("/metrics", monitor.New())
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Grouping the routes into one prefix
	v1 := app.Group("/api/v1")
}