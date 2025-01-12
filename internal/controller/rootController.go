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

	r.Get("/FindCustomerByID", handler.CustomerService.FindCustomerByID)
	r.Post("/CreateCustomer", handler.CustomerService.CreateCustomer)
	r.Delete("/DeleteCustomerByID", handler.CustomerService.DeleteCustomerByID)
	r.Patch("/UpdateCustomer", handler.CustomerService.UpdateTableNumberInCustomer)

	//r.Post("/FindOrderLogWithFK", handler.FindOrderLogWithFK)

	r.Patch("/MakePayment", handler.CustomerService.MakePayment)
	r.Post("/CreateOrderLog", handler.OrderLogService.CreateOrderLog)
	r.Post("/CreateRestaurant", handler.RestaurantService.CreateRestaurant)
	r.Post("/CreateMenu", handler.MenuService.CreateMenu)
	r.Post("/CreateMenuItem", handler.MenuItemService.CreateMenuItem)
	r.Patch("/UpdateMenuItem", handler.MenuItemService.UpdateMenuItem)

	return r
}
