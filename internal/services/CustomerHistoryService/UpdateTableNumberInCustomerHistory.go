package CustomerHistoryService

import (
	"food-order/internal/config"
	"food-order/internal/repositories"
	"food-order/internal/utils"

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
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryID) {
		return utils.SendBadRequest(ctx, &response, "id is not exist")
	}

	targetCustomerHistory, err := ch.CustomerHistoryRepository.FindOneById(ctx.Context(), request.CustomerHistoryID)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	oldTable, err := tableInfoRepository.FindOneByTableNumber(ctx.Context(), targetCustomerHistory.TableNumber)
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
	targetCustomerHistory.TableNumber = request.TableNumber
	if err := ch.CustomerHistoryRepository.UpdateOne(ctx.Context(), targetCustomerHistory); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["message"] = "Updated"
	return ctx.JSON(response)

}
