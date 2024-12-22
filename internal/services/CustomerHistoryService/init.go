package CustomerHistoryService

import (
	"fmt"
	"food-order/internal/models"
	"food-order/internal/repositories"
	"reflect"
)

type CustomerHistoryService struct {
	CustomerHistoryRepository *repositories.CustomerHistoryRepository
}

func GetCustomerHistoryService() *CustomerHistoryService {
	return &CustomerHistoryService{
		CustomerHistoryRepository: repositories.GetCustomerHistoryRepository(),
	}
}

func CustomerHistoryObjectToString(obj *models.CustomerHistory) string {
	result := map[string]string{}

	v := reflect.ValueOf(obj).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		fieldName := t.Field(i).Name
		fieldValue := v.Field(i).Interface()
		if fieldName == "BaseModel" {
			continue
		}
		result[fieldName] = fmt.Sprintf("%v", fieldValue)
	}

	return fmt.Sprintf("%+v", result)
}
