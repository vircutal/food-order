package repositories

import "food-order/internal/models"

type MenuItemRepository struct {
	*BaseDB[models.MenuItem]
}

func GetMenuItemRepository() *MenuItemRepository {
	return &MenuItemRepository{
		BaseDB: (GetBaseDB[models.MenuItem]()),
	}
}
