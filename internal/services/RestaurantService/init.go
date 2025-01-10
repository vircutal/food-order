package RestaurantService

import "food-order/internal/repositories"

type RestaurantService struct {
	RestaurantRepository *repositories.RestaurantRepository
}

func GetRestaurantService() *RestaurantService {
	return &RestaurantService{
		RestaurantRepository: repositories.GetRestaurantRepository(),
	}
}
