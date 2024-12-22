package MenuService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteMenuRequest struct {
	MenuID uuid.UUID `json:"menu_id"`
}

func (m *MenuService) DeleteMenu(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request DeleteMenuRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************
	if !m.MenuRepository.CheckExistByID(ctx.Context(), request.MenuID) {
		return utils.SendBadRequest(ctx, &response, "menu does not exist")
	}

	if err := m.MenuRepository.DeleteOneById(ctx.Context(), request.MenuID); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["message"] = "deleted"
	return ctx.JSON(response)
}
