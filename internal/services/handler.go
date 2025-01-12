package services

import (
	"food-order/internal/services/CustomerService"
	"food-order/internal/services/MenuItemService"
	"food-order/internal/services/MenuService"
	"food-order/internal/services/OrderLogService"
	"food-order/internal/services/RestaurantService"
	"food-order/internal/services/TableInfoService"
)

type Handler struct {
	CustomerService   *CustomerService.CustomerService
	OrderLogService   *OrderLogService.OrderLogService
	TableInfoService  *TableInfoService.TableInfoService
	RestaurantService *RestaurantService.RestaurantService
	MenuService       *MenuService.MenuService
	MenuItemService   *MenuItemService.MenuItemService
}

func GetHandler() *Handler {
	return &Handler{
		CustomerService:   CustomerService.GetCustomerService(),
		OrderLogService:   OrderLogService.GetOrderLogService(),
		TableInfoService:  TableInfoService.GetTableInfoService(),
		RestaurantService: RestaurantService.GetRestaurantService(),
		MenuService:       MenuService.GetMenuService(),
		MenuItemService:   MenuItemService.GetMenuItemService(),
	}
}
