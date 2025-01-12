package repositories

import (
	"food-order/internal/models"
)

type CustomerRepository struct {
	*BaseDB[models.Customer]
}

func GetCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		BaseDB: GetBaseDB[models.Customer](),
	}
}
