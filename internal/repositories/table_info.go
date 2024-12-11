package repositories

import "food-order/internal/models"

type TableInfoRepository struct {
	*BaseDB[models.TableInfo]
}

func GetTableInfoRepository() *TableInfoRepository {
	return &TableInfoRepository{
		BaseDB: GetBaseDB[models.TableInfo](),
	}
}
