package CustomerService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FindCustomerByIDRequest struct {
	CustomerID uuid.UUID `json:"customer_id"`
}

// Not done
func (ch *CustomerService) FindCustomerByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request FindCustomerByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendBadRequest(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ch.CustomerRepository.CheckExistByID(ctx.Context(), request.CustomerID) {
		return utils.SendBadRequest(ctx, &response, "Customer is not exist")
	}

	targetCustomer, err := ch.CustomerRepository.FindOneById(ctx.Context(), request.CustomerID)
	if err != nil {
		return utils.SendBadRequest(ctx, &response, err.Error())
	}

	response["result"] = *targetCustomer
	return ctx.JSON(response)
}
