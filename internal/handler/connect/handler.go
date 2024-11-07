package connect

import "github.com/gofiber/fiber/v2"

type Handler interface {
	doGetHalloWorld(c *fiber.Ctx) error
}

type impl struct {
	manager 
}