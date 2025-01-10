package RestaurantService

import (
	"food-order/internal/models"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateReataurantRequest struct {
	RestaurantName string `json:"restaurant_name"`
	Branch         string `json:"branch"`
}

func (rs *RestaurantService) CreateRestaurant(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateReataurantRequest
	response := map[string]interface{}{}
	//**************************************************************
	//**************************************************************

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	targetRestaurant := &models.Restaurant{
		ID:             uuid.New(),
		RestaurantName: request.RestaurantName,
		Branch:         request.Branch,
	}

	if err := rs.RestaurantRepository.AddOne(ctx.Context(), targetRestaurant); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetRestaurant
	return ctx.Status(201).JSON(response)
}
