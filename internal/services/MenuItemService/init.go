package MenuItemService

import "food-order/internal/repositories"

type MenuItemService struct {
	MenuItemRepository *repositories.MenuItemRepository
}

func GetMenuItemService() *MenuItemService {
	return &MenuItemService{
		MenuItemRepository: repositories.GetMenuItemRepository(),
	}
}
