package OrderLogService

import (
	"food-order/internal/models"
	"food-order/internal/services/CustomerHistoryService"
	"food-order/internal/services/MenuService"
	"food-order/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateOrderLogRequest struct {
	CustomerID       uuid.UUID `json:"customer_id"`
	FoodID           uuid.UUID `json:"food_id"`
	FoodPrice        string    `json:"food_price"`
	Quantity         string    `json:"quantity"`
	OrderDescription string    `json:"order_description"`
}

func (ol *OrderLogService) CreateOrderLog(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateOrderLogRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	foodPrice, _ := strconv.ParseFloat(request.FoodPrice, 64)
	quantity, _ := strconv.Atoi(request.Quantity)

	customerHistoryService := CustomerHistoryService.GetCustomerHistoryService()
	menuService := MenuService.GetMenuService()

	//**************************************************************
	//**************************************************************

	//Check if customer id exist
	if !customerHistoryService.CustomerHistoryRepository.CheckExistByID(ctx.Context(), request.CustomerID) {
		return utils.SendBadRequest(ctx, &response, "customer is not exist")
	}

	//Check if food id exist
	if !menuService.MenuRepository.CheckExistByID(ctx.Context(), request.FoodID) {
		return utils.SendBadRequest(ctx, &response, "this menu is not exist")
	}

	newOrder := &models.OrderLog{
		ID:               uuid.New(),
		CustomerID:       request.CustomerID,
		FoodID:           request.FoodID,
		FoodPrice:        foodPrice,
		Quantity:         quantity,
		OrderDescription: request.OrderDescription,
		OrderedTime:      time.Now(),
	}

	if err := ol.OrderLogRepository.AddOne(ctx.Context(), newOrder); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	return nil
}
