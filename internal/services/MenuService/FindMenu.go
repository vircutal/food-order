package MenuService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type FindMenuByIDRequest struct {
	MenuID uuid.UUID `json:"menu_id"`
}

func (m *MenuService) FindMenuByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request FindMenuByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************
	if !m.MenuRepository.CheckExistByID(ctx.Context(), request.MenuID) {
		return utils.SendBadRequest(ctx, &response, "menu does not exist")
	}

	targetMenu, err := m.MenuRepository.FindOneById(ctx.Context(), request.MenuID)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["result"] = *targetMenu
	return ctx.JSON(response)
}
