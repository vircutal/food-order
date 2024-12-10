package repositories

import "food-order/internal/models"

type MenuRepository struct {
	*BaseDB[models.Menu]
}

func GetMenuRepository() *MenuRepository {
	return &MenuRepository{
		BaseDB: GetBaseDB[models.Menu](),
	}
}
