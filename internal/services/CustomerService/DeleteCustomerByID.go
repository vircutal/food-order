package CustomerService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteCustomerByIDRequest struct {
	CustomerID uuid.UUID `json:"customer_id"`
}

// Need to handle delete function properly
func (ch *CustomerService) DeleteCustomerByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request DeleteCustomerByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ch.CustomerRepository.CheckExistByID(ctx.Context(), request.CustomerID) {
		return utils.SendBadRequest(ctx, &response, "Table is not exist")
	}

	if err := ch.CustomerRepository.DeleteOneById(ctx.Context(), request.CustomerID); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	response["message"] = "Deleted"
	response["id"] = request.CustomerID.String()
	return ctx.JSON(response)
}
