package repositories

import "food-order/internal/models"

type OrderLogRepository struct {
	*BaseDB[models.OrderLog]
}

func GetOrderLogRepository() *OrderLogRepository {
	return &OrderLogRepository{
		BaseDB: GetBaseDB[models.OrderLog](),
	}
}
