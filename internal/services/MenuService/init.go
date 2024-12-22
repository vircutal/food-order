package MenuService

import "food-order/internal/repositories"

type MenuService struct {
	MenuRepository *repositories.MenuRepository
}

func GetMenuService() *MenuService {
	return &MenuService{
		MenuRepository: repositories.GetMenuRepository(),
	}
}
