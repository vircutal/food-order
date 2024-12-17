package controller

import (
	"food-order/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetRootController() *fiber.App {
	r := fiber.New()
	handler := services.GetHandler()
	// all the handling paths are here
	// handler function should be specified in services folder
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	r.Get("/test", handler.TableInfoService.GetTableInfoByStatus)
	r.Patch("/test", handler.TableInfoService.UpdateTableInfo)

	//---------------------------------

	return r
}
