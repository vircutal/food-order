package CustomerService

import (
	"food-order/internal/config"
	"food-order/internal/repositories"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateTableNumberInCustomerRequest struct {
	CustomerID  uuid.UUID `json:"customer_id"`
	TableNumber int       `json:"table_number"`
}

func (ch *CustomerService) UpdateTableNumberInCustomer(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateTableNumberInCustomerRequest
	response := map[string]interface{}{}
	tableInfoRepository := repositories.GetTableInfoRepository()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !ch.CustomerRepository.CheckExistByID(ctx.Context(), request.CustomerID) {
		return utils.SendBadRequest(ctx, &response, "id is not exist")
	}

	targetCustomer, err := ch.CustomerRepository.FindOneById(ctx.Context(), request.CustomerID)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	oldTable, err := tableInfoRepository.FindOneByTableNumber(ctx.Context(), targetCustomer.TableNumber)
	if err := tableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	newTable, err := tableInfoRepository.FindOneByTableNumber(ctx.Context(), request.TableNumber)
	if err := tableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if newTable.Status != config.TableIsAvailable {
		return utils.SendBadRequest(ctx, &response, "table is not available")
	}
	//**************************************************************
	//**************************************************************

	//release old table
	oldTable.Status = config.TableIsAvailable
	if err := tableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	// lock new table
	newTable.Status = config.TableIsOccupied
	if err := tableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	//update customer
	targetCustomer.TableNumber = request.TableNumber
	if err := ch.CustomerRepository.UpdateOne(ctx.Context(), targetCustomer); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["message"] = "Updated"
	return ctx.Status(200).JSON(response)

}
