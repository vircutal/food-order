package CustomerService

import (
	"food-order/internal/config"
	"food-order/internal/models"
	"food-order/internal/repositories"
	"food-order/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateCustomerRequest struct {
	TableNumber string `json:"table_number"`
	Status      string `json:"status"`
}

func (ch *CustomerService) CreateCustomer(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateCustomerRequest
	response := map[string]interface{}{}

	//Get service
	tableInfoRepository := repositories.GetTableInfoRepository()

	//parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//convert string to integer
	TableNumber, _ := strconv.Atoi(request.TableNumber)
	//**************************************************************
	//**************************************************************

	//check if table_number exists
	if !tableInfoRepository.CheckTableNumberExist(ctx.Context(), TableNumber) {
		return utils.SendBadRequest(ctx, &response, "Table doesnt exist")
	}

	//check if table is available
	if !tableInfoRepository.CheckTableStatus(ctx.Context(), TableNumber, config.TableIsAvailable) {
		return utils.SendBadRequest(ctx, &response, "Table is not available")
	}

	//check if request's status is valid

	if !config.CustomerStatus[request.Status] {
		return utils.SendBadRequest(ctx, &response, "Status is invalid")
	}

	//Add data to db
	targetCustomer := &models.Customer{
		ID:          uuid.New(),
		TableNumber: TableNumber,
		Status:      request.Status,
		TimeIn:      time.Now(),
		TimeOut:     nil,
		PaymentTime: nil,
		TotalPrice:  nil,
	}

	err := ch.CustomerRepository.AddOne(ctx.Context(), targetCustomer)

	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	} else {
		return utils.SendStatusOK(ctx, &response, "New Customer Added")
	}

}
