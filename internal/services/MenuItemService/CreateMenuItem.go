package MenuItemService

import (
	"food-order/internal/models"
	"food-order/internal/repositories"
	"food-order/internal/utils"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CreateMenuItemRequest struct {
	MenuID              uuid.UUID            `form:"menu_id"`
	MenuItemName        string               `form:"menu_item_name"`
	MenuItemPrice       float64              `form:"men_item_price"`
	MenuItemDescription string               `form:"menu_item_description"`
	MenuItemImageURL    multipart.FileHeader `form:"menu_item_image_url"`
}

func (mi *MenuItemService) CreateMenuItem(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateMenuItemRequest
	response := map[string]interface{}{}

	menuRepository := repositories.GetMenuRepository()
	S3Client := utils.GetS3Client()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	file2, _ := request.MenuItemImageURL.Open()

	if !menuRepository.CheckExistByID(ctx.Context(), request.MenuID) {
		return utils.SendBadRequest(ctx, &response, "MenuID is not exist")
	}
	//**************************************************************
	//**************************************************************

	if err := S3Client.S3PutObject(ctx.Context(), "my-bucket", request.MenuItemImageURL.Filename, file2); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	targetMenuItem := &models.MenuItem{
		ID:                  uuid.New(),
		MenuID:              request.MenuID,
		MenuItemName:        request.MenuItemName,
		MenuItemPrice:       request.MenuItemPrice,
		MenuItemDescription: &request.MenuItemDescription,
		MenuItemImageURL:    &request.MenuItemImageURL.Filename,
	}

	if err := mi.MenuItemRepository.AddOne(ctx.Context(), targetMenuItem); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetMenuItem
	return ctx.Status(201).JSON(response)
}
