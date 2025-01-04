package repositories

import "food-order/internal/models"

type RestaurantRepository struct {
	*BaseDB[models.Restaurant]
}

func GetRestaurantRepository() *RestaurantRepository {
	return &RestaurantRepository{
		BaseDB: (GetBaseDB[models.Restaurant]()),
	}
}
