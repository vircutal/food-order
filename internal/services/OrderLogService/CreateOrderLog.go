package OrderLogService

import (
	"food-order/internal/models"
	"food-order/internal/repositories"
	"food-order/internal/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateOrderLogRequest struct {
	MenuItemID       uuid.UUID `json:"menu_item_id"`
	CustomerID       uuid.UUID `json:"customer_id"`
	MenuItemPrice    string    `json:"menu_item_price"`
	Quantity         string    `json:"quantity"`
	OrderDescription string    `json:"order_description"`
}

func (orderLogService *OrderLogService) CreateOrderLog(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateOrderLogRequest
	response := map[string]interface{}{}

	menuItemRepository := repositories.GetMenuItemRepository()
	customerRepository := repositories.GetCustomerRepository()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	//check if menu_item_id exists
	if !menuItemRepository.CheckExistByID(ctx.Context(), request.MenuItemID) {
		return utils.SendBadRequest(ctx, &response, "menu_item_id is not exist")
	}

	//check if customer__id exists
	if !customerRepository.CheckExistByID(ctx.Context(), request.CustomerID) {
		return utils.SendBadRequest(ctx, &response, "customer_id is not exist")
	}

	quantity, _ := strconv.Atoi(request.Quantity)
	menuItemPrice, _ := strconv.ParseFloat(request.MenuItemPrice, 64)

	//**************************************************************
	//**************************************************************

	newOrderLog := models.OrderLog{
		ID:               uuid.New(),
		CustomerID:       request.CustomerID,
		MenuItemID:       request.MenuItemID,
		MenuItemPrice:    menuItemPrice,
		Quantity:         quantity,
		OrderedTime:      time.Now(),
		OrderDescription: &request.OrderDescription,
	}

	if err := orderLogService.OrderLogRepository.AddOne(ctx.Context(), &newOrderLog); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	ctx.Status(200)
	response["result"] = newOrderLog
	return ctx.JSON(response)
}
