package repositories

import (
	"context"
	"food-order/internal/models"
)

type TableInfoRepository struct {
	*BaseDB[models.TableInfo]
}

func GetTableInfoRepository() *TableInfoRepository {
	return &TableInfoRepository{
		BaseDB: GetBaseDB[models.TableInfo](),
	}
}

func (t *TableInfoRepository) GetAllByStatus(ctx context.Context, status string) (*[]models.TableInfo, error) {
	var models []models.TableInfo
	err := t.db.NewSelect().Model(&models).Where("status = ?", status).Scan(ctx)
	return &models, err
}
