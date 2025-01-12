package CustomerService

import (
	"food-order/internal/repositories"
)

type CustomerService struct {
	CustomerRepository *repositories.CustomerRepository
}

func GetCustomerService() *CustomerService {
	return &CustomerService{
		CustomerRepository: repositories.GetCustomerRepository(),
	}
}
