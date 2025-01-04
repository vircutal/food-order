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
	r.Get("/FindAllTableInfoByStatus", handler.TableInfoService.FindAllTableInfoByStatus)
	r.Patch("/UpdateTableInfo", handler.TableInfoService.UpdateTableInfoStatus)
	r.Post("/CreateTableInfo", handler.TableInfoService.CreateTableInfo)
	r.Delete("/DeleteTableInfoByID", handler.TableInfoService.DeleteTableInfoByID)

	r.Get("/FindCustomerHistoryByID", handler.CustomerHistoryService.FindCustomerHistoryByID)
	r.Post("/CreateCustomerHistory", handler.CustomerHistoryService.CreateCustomerHistory)
	r.Delete("/DeleteCustomerHistoryByID", handler.CustomerHistoryService.DeleteCustomerHistoryByID)
	r.Patch("/UpdateCustomerHistory", handler.CustomerHistoryService.UpdateTableNumberInCustomerHistory)

	//r.Post("/FindOrderLogWithFK", handler.FindOrderLogWithFK)

	r.Patch("/MakePayment", handler.CustomerHistoryService.MakePayment)

	return r
}
