package repositories

import "food-order/internal/models"

type CustomerHistoryRepository struct {
	*BaseDB[models.CustomerHistory]
}

func GetCustomerHistoryRepository() *CustomerHistoryRepository {
	return &CustomerHistoryRepository{
		BaseDB: GetBaseDB[models.CustomerHistory](),
	}
}
