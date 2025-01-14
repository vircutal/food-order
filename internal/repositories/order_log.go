package repositories

import (
	"context"
	"food-order/internal/models"

	"github.com/google/uuid"
)

type OrderLogRepository struct {
	*BaseDB[models.OrderLog]
}

func GetOrderLogRepository() *OrderLogRepository {
	return &OrderLogRepository{
		BaseDB: GetBaseDB[models.OrderLog](),
	}
}

func (ol *OrderLogRepository) CheckOrderExistByIDs(ctx context.Context, customer_id, food_id uuid.UUID) bool {
	var model models.OrderLog
	err := ol.db.NewSelect().Model(&model).Where("customer_id = ?", customer_id).Where("food_id = ?", food_id).Scan(ctx)

	if err != nil {
		return false
	}

	return true
}

func (ol *OrderLogRepository) FindAllByCustomerID(ctx context.Context, customer_id uuid.UUID) (*[]models.OrderLog, error) {
	var model []models.OrderLog
	err := ol.db.NewSelect().Model(&model).Where("customer_id = ?", customer_id).Scan(ctx)
	return &model, err
}
