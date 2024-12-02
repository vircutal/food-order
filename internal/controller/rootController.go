package controller

import (
	"github.com/gofiber/fiber/v2"
)

func GetRootController() *fiber.App {
	r := fiber.New()
	// all the handling paths are here
	// handler function should be specified in services folder
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//---------------------------------

	return r
}
