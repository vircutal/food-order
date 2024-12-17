package services

import (
	"food-order/internal/services/CustomerHistoryService"
	"food-order/internal/services/MenuService"
	"food-order/internal/services/OrderLogService"
	"food-order/internal/services/TableInfoService"
)

type Handler struct {
	*CustomerHistoryService.CustomerHistoryService
	*MenuService.MenuService
	*OrderLogService.OrderLogService
	*TableInfoService.TableInfoService
}

func GetHandler() *Handler {
	return &Handler{
		CustomerHistoryService: CustomerHistoryService.GetCustomerHistoryService(),
		MenuService:            MenuService.GetMenuService(),
		OrderLogService:        OrderLogService.GetOrderLogService(),
		TableInfoService:       TableInfoService.GetTableInfoService(),
	}
}
