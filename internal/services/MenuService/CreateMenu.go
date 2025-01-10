package MenuService

import (
	"food-order/internal/models"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	MenuName     string    `json:"menu_name"`
	RestaurantID uuid.UUID `json:"reataurant_id"`
}

func (ms *MenuService) CreateMenu(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateMenuRequest
	response := map[string]interface{}{}
	//**************************************************************
	//**************************************************************

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	targetMenu := &models.Menu{
		ID:           uuid.New(),
		MenuName:     request.MenuName,
		RestaurantID: request.RestaurantID,
	}

	if err := ms.MenuRepository.AddOne(ctx.Context(), targetMenu); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetMenu
	return ctx.Status(201).JSON(response)

}
