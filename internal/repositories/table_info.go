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

func (t *TableInfoRepository) GetTableInfoByTableNumber(ctx context.Context, TableNumber int) (*models.TableInfo, error) {
	var model models.TableInfo
	err := t.db.NewSelect().Model(&model).Where("table_number = ?", TableNumber).Scan(ctx)
	return &model, err
}

func (t *TableInfoRepository) CheckTableNumberExist(ctx context.Context, TableNumber int) bool {
	var model models.TableInfo
	err := t.db.NewSelect().Model(&model).Where("table_number = ?", TableNumber).Scan(ctx)
	if err == nil {
		return true
	}
	return false
}
