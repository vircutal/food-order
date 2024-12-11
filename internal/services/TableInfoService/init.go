package TableInfoService

import "food-order/internal/repositories"

type TableInfoService struct {
	TableInfoRepository *repositories.TableInfoRepository
}

func GetTableInfoService() *TableInfoService {
	return &TableInfoService{
		TableInfoRepository: repositories.GetTableInfoRepository(),
	}
}
