package CustomerService

import (
	"food-order/internal/repositories"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type MakePaymentRequest struct {
	CustomerId uuid.UUID `json:"customer_id"`
}

func (ch *CustomerService) MakePayment(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request MakePaymentRequest
	response := map[string]interface{}{}
	orderLogRepository := repositories.GetOrderLogRepository()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !ch.CustomerRepository.CheckExistByID(ctx.Context(), request.CustomerId) {
		return utils.SendBadRequest(ctx, &response, "Customer id is not exist")
	}

	orders, err := orderLogRepository.FindAllByCustomerID(ctx.Context(), request.CustomerId)

	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	//**************************************************************
	//**************************************************************

	totalPrice := 0.00
	for _, val := range *orders {
		if val.MenuItemPrice < 0 {
			return utils.SendInternalServerError(ctx, &response, "Food Price is a negative value. Need Admin's check")
		}
		totalPrice += (val.MenuItemPrice * float64(val.Quantity))
	}

	targetCustomer, err := ch.CustomerRepository.FindOneById(ctx.Context(), request.CustomerId)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	targetCustomer.TotalPrice = &totalPrice

	if err := ch.CustomerRepository.UpdateOne(ctx.Context(), targetCustomer); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetCustomer
	return ctx.JSON(response)
}
