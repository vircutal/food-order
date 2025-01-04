package CustomerHistoryService

import (
	"food-order/internal/repositories"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MakePaymentRequest struct {
	CustomerHistoryId uuid.UUID `json:"customer_history_id"`
}

func (ch *CustomerHistoryService) MakePayment(ctx *fiber.Ctx) error {
	var request MakePaymentRequest
	response := map[string]interface{}{}
	orderLogRepository := repositories.GetOrderLogRepository()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !ch.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerHistoryId) {
		return utils.SendBadRequest(ctx, &response, "Customer id is not exist")
	}

	orders, err := orderLogRepository.FindAllByCustomerHistoryID(ctx.Context(), request.CustomerHistoryId)

	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	totalPrice := 0.00
	for _, val := range *orders {
		if val.FoodPrice < 0 {
			return utils.SendInternalServerError(ctx, &response, "Food Price is a negative value. Need Admin's check")
		}
		totalPrice += (val.FoodPrice * float64(val.Quantity))
	}

	//fmt.Println(totalPrice)

	targetCustomerHistory, err := ch.CustomerHistoryRepository.FindOneById(ctx.Context(), request.CustomerHistoryId)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if targetCustomerHistory.TotalPrice == nil {
		targetCustomerHistory.TotalPrice = &totalPrice
	}

	if err := ch.CustomerHistoryRepository.UpdateOne(ctx.Context(), targetCustomerHistory); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetCustomerHistory
	return ctx.JSON(response)
}
