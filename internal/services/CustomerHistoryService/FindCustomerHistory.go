package CustomerHistoryService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FindCustomerHistoryByIDRequest struct {
	CustomerHistoryID uuid.UUID `json:"customer_history_id"`
}

// Not done
func (ch *CustomerHistoryService) FindCustomerHistoryByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request FindCustomerHistoryByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendBadRequest(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryID) {
		return utils.SendBadRequest(ctx, &response, "Customer History is not exist")
	}

	targetCustomerHistory, err := ch.CustomerHistoryRepository.FindOneById(ctx.Context(), request.CustomerHistoryID)
	if err != nil {
		return utils.SendBadRequest(ctx, &response, err.Error())
	}

	response["result"] = *targetCustomerHistory
	return ctx.JSON(response)
}
