package CustomerHistoryService

import (
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
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}
	//**************************************************************
	//**************************************************************

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryID) {
		ctx.Status(400)
		response["message"] = "Customer History is not exist"
		return ctx.JSON(response)
	}

	targetCustomerHistory, err := ch.CustomerHistoryRepository.FindOneById(ctx.Context(), request.CustomerHistoryID)
	if err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	response["result"] = *targetCustomerHistory
	return ctx.JSON(response)
}
