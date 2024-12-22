package CustomerHistoryService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteCustomerHistoryByIDRequest struct {
	CustomerHistoryID uuid.UUID `json:"customer_history_id"`
}

func (ch *CustomerHistoryService) DeleteCustomerHistoryByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request DeleteCustomerHistoryByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryID) {
		return utils.SendBadRequest(ctx, &response, "Table is not exist")
	}

	if err := ch.CustomerHistoryRepository.DeleteOneById(ctx.Context(), request.CustomerHistoryID); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	response["message"] = "Deleted"
	response["id"] = request.CustomerHistoryID.String()
	return ctx.JSON(response)
}
