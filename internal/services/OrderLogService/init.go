package OrderLogService

import "food-order/internal/repositories"

type OrderLogService struct {
	OrderLogRepository *repositories.OrderLogRepository
}

func GetOrderLogService() *OrderLogService {
	return &OrderLogService{
		OrderLogRepository: repositories.GetOrderLogRepository(),
	}
}
