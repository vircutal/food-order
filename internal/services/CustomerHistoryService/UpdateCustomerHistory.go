package CustomerHistoryService

import (
	"food-order/internal/config"
	"food-order/internal/services/TableInfoService"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateCustomerHistoryRequest struct {
	CustomerHistoryID uuid.UUID `json:"customer_history_id"`
	TableNumber       int       `json:"table_number"`
	//TimeOut           *time.Time `json:"time_out"`
}

func (ch *CustomerHistoryService) UpdateCustomerHistory(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateCustomerHistoryRequest
	response := map[string]interface{}{}
	tableInfoService := TableInfoService.GetTableInfoService()

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

	oldTable, err := tableInfoService.TableInfoRepository.FindOneByTableNumber(ctx.Context(), targetCustomerHistory.TableNumber)
	if err := tableInfoService.TableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}
	newTable, err := tableInfoService.TableInfoRepository.FindOneByTableNumber(ctx.Context(), request.TableNumber)
	if err := tableInfoService.TableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
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
	if err := tableInfoService.TableInfoRepository.UpdateOne(ctx.Context(), oldTable); err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	// lock new table
	newTable.Status = config.TableIsOccupied
	if err := tableInfoService.TableInfoRepository.UpdateOne(ctx.Context(), newTable); err != nil {
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
