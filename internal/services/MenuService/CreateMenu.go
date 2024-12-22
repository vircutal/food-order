package MenuService

import (
	"food-order/internal/models"
	"food-order/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateMenuRequest struct {
	FoodName        string  `json:"food_name"`
	FoodPrice       string  `json:"food_price"`
	FoodDescription *string `json:"food_description"`
}

func (m *MenuService) CreateMenu(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateMenuRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	//convert string to float
	foodPrice, _ := strconv.ParseFloat(request.FoodPrice, 64)

	//init object
	targetMenu := &models.Menu{
		ID:              uuid.New(),
		FoodName:        request.FoodName,
		FoodPrice:       foodPrice,
		FoodDescription: request.FoodDescription,
		FoodImageURL:    nil,
	}

	if err := m.MenuRepository.AddOne(ctx.Context(), targetMenu); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["message"] = "Added"
	response["result"] = *targetMenu
	return ctx.JSON(response)

}
