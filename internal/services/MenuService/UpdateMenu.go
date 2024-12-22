package MenuService

import (
	"food-order/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateMenuByIDRequest struct {
	MenuID          uuid.UUID `json:"menu_id"`
	FoodName        *string   `json:"food_name"`
	FoodPrice       *string   `json:"food_price"`
	FoodDescription *string   `json:"food_description"`
}

func (m *MenuService) UpdateMenuByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateMenuByIDRequest
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

	if request.FoodName != nil {
		targetMenu.FoodName = *request.FoodName
	}
	if request.FoodDescription != nil {
		targetMenu.FoodDescription = request.FoodDescription
	}
	if request.FoodPrice != nil {
		//convert string to float
		foodPrice, _ := strconv.ParseFloat(*request.FoodPrice, 64)
		targetMenu.FoodPrice = foodPrice
	}
	ctx.Status(200)
	response["result"] = *targetMenu
	return ctx.JSON(response)
}
