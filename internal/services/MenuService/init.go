package MenuService

import "food-order/internal/repositories"

type MenuService struct {
	UserRepository *repositories.MenuRepository
}

func GetMenuService() *MenuService {
	return &MenuService{
		UserRepository: repositories.GetMenuRepository(),
	}
}
