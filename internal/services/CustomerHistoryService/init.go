package CustomerHistoryService

import (
	"food-order/internal/repositories"
)

type CustomerHistoryService struct {
	CustomerHistoryRepository *repositories.CustomerHistoryRepository
}

func GetCustomerHistoryService() *CustomerHistoryService {
	return &CustomerHistoryService{
		CustomerHistoryRepository: repositories.GetCustomerHistoryRepository(),
	}
}
