package CustomerHistoryService

import (
	"food-order/internal/config"
	"food-order/internal/models"
	"food-order/internal/services/TableInfoService"
	"food-order/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateCustomerHistoryRequest struct {
	TableNumber string `json:"table_number"`
	Status      string `json:"status"`
}

func (ch *CustomerHistoryService) CreateCustomerHistory(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateCustomerHistoryRequest
	response := map[string]interface{}{}

	//Get service
	tableInfoService := TableInfoService.GetTableInfoService()

	//parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//convert string to integer
	TableNumber, _ := strconv.Atoi(request.TableNumber)
	//**************************************************************
	//**************************************************************

	//check if table_number exists
	if !tableInfoService.TableInfoRepository.CheckTableNumberExist(ctx.Context(), TableNumber) {
		return utils.SendBadRequest(ctx, &response, "Table doesnt exist")
	}

	//check if table is available
	if !tableInfoService.TableInfoRepository.CheckTableStatus(ctx.Context(), TableNumber, config.TableIsAvailable) {
		return utils.SendBadRequest(ctx, &response, "Table is not available")
	}

	//check if request's status is valid

	if !config.CustomerHistoryStatus[request.Status] {
		return utils.SendBadRequest(ctx, &response, "Status is invalid")
	}

	//Add data to db
	targetCustomerHistory := &models.CustomerHistory{
		ID:          uuid.New(),
		TableNumber: TableNumber,
		Status:      request.Status,
		TimeIn:      time.Now(),
		TimeOut:     nil,
		PaymentTime: nil,
		TotalPrice:  nil,
	}

	err := ch.CustomerHistoryRepository.AddOne(ctx.Context(), targetCustomerHistory)

	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	} else {
		return utils.SendStatusOK(ctx, &response, "New Customer History Added")
	}

}
