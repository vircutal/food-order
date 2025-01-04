package services

import (
	"food-order/internal/services/CustomerHistoryService"
	"food-order/internal/services/OrderLogService"
	"food-order/internal/services/TableInfoService"
)

type Handler struct {
	CustomerHistoryService *CustomerHistoryService.CustomerHistoryService
	OrderLogService        *OrderLogService.OrderLogService
	TableInfoService       *TableInfoService.TableInfoService
}

func GetHandler() *Handler {
	return &Handler{
		CustomerHistoryService: CustomerHistoryService.GetCustomerHistoryService(),
		OrderLogService:        OrderLogService.GetOrderLogService(),
		TableInfoService:       TableInfoService.GetTableInfoService(),
	}
}
