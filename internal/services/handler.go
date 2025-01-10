package services

import (
	"food-order/internal/services/CustomerHistoryService"
	"food-order/internal/services/MenuItemService"
	"food-order/internal/services/MenuService"
	"food-order/internal/services/OrderLogService"
	"food-order/internal/services/RestaurantService"
	"food-order/internal/services/TableInfoService"
)

type Handler struct {
	CustomerHistoryService *CustomerHistoryService.CustomerHistoryService
	OrderLogService        *OrderLogService.OrderLogService
	TableInfoService       *TableInfoService.TableInfoService
	RestaurantService      *RestaurantService.RestaurantService
	MenuService            *MenuService.MenuService
	MenuItemService        *MenuItemService.MenuItemService
}

func GetHandler() *Handler {
	return &Handler{
		CustomerHistoryService: CustomerHistoryService.GetCustomerHistoryService(),
		OrderLogService:        OrderLogService.GetOrderLogService(),
		TableInfoService:       TableInfoService.GetTableInfoService(),
		RestaurantService:      RestaurantService.GetRestaurantService(),
		MenuService:            MenuService.GetMenuService(),
		MenuItemService:        MenuItemService.GetMenuItemService(),
	}
}
