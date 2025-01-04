package CustomerHistoryService

import (
	"food-order/internal/config"
	"food-order/internal/repositories"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateTableNumberInCustomerHistoryRequest struct {
	CustomerHistoryID uuid.UUID `json:"customer_history_id"`
	TableNumber       int       `json:"table_number"`
}

func (ch *CustomerHistoryService) UpdateTableNumberInCustomerHistory(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateTableNumberInCustomerHistoryRequest
	response := map[string]interface{}{}
	tableInfoRepository := repositories.GetTableInfoRepository()

	if err := ctx.BodyParser(&request); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryID) {
		ctx.Status(400)
		response["message"] = "id is not exist"
		return ctx.JSON(response)
	}

	targetCustomerHistory, err := ch.CustomerHistoryRepository.FindOneById(ctx.Context(), request.CustomerHistoryID)
	if err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	oldTable, err := tableInfoRepository.FindOneByTableNumber(ctx.Context(), targetCustomerHistory.TableNumber)
	if err := tableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}
	newTable, err := tableInfoRepository.FindOneByTableNumber(ctx.Context(), request.TableNumber)
	if err := tableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	if newTable.Status != config.TableIsAvailable {
		ctx.Status(400)
		response["message"] = "table is not available"
		return ctx.JSON(response)
	}
	//**************************************************************
	//**************************************************************

	//release old table
	oldTable.Status = config.TableIsAvailable
	if err := tableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	// lock new table
	newTable.Status = config.TableIsOccupied
	if err := tableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	//update customer
	targetCustomerHistory.TableNumber = request.TableNumber
	if err := ch.CustomerHistoryRepository.UpdateOne(ctx.Context(), targetCustomerHistory); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	ctx.Status(200)
	response["message"] = "Updated"
	return ctx.JSON(response)

}
