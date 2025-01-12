package MenuItemService

import (
	"food-order/internal/utils"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateMenuItemRequest struct {
	MenuItemID          uuid.UUID             `form:"menu_item_id"`
	MenuItemName        *string               `form:"menu_item_name"`
	MenuItemPrice       *float64              `form:"menu_item_price"`
	MenuItemDescription *string               `form:"menu_item_description"`
	MenuItemImage       *multipart.FileHeader `form:"menu_item_image"`
}

func (mi *MenuItemService) UpdateMenuItem(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateMenuItemRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !mi.MenuItemRepository.CheckExistByID(ctx.Context(), request.MenuItemID) {
		return utils.SendBadRequest(ctx, &response, "Menu item id is not exist")
	}

	targetMenuItem, err := mi.MenuItemRepository.FindOneById(ctx.Context(), request.MenuItemID)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	//**************************************************************
	//**************************************************************

	menuItemImageFileHeader, _ := ctx.FormFile("menu_item_image")
	if menuItemImageFileHeader != nil {
		menuItemImage, _ := menuItemImageFileHeader.Open()
		storage := utils.GetStorageClient()
		key := utils.GenerateKeyFromFilename(menuItemImageFileHeader.Filename, request.MenuItemID.String())
		if err := storage.S3PutObject(ctx.Context(), "my-bucket", *key, menuItemImage); err != nil {
			return utils.SendInternalServerError(ctx, &response, err.Error())
		}
	}
	if request.MenuItemName != nil {
		targetMenuItem.MenuItemName = *request.MenuItemName
	}
	if request.MenuItemPrice != nil {
		targetMenuItem.MenuItemPrice = *request.MenuItemPrice
	}
	if request.MenuItemDescription != nil {
		targetMenuItem.MenuItemDescription = request.MenuItemDescription
	}

	if err := mi.MenuItemRepository.UpdateOne(ctx.Context(), targetMenuItem); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	response["result"] = *targetMenuItem
	return ctx.Status(200).JSON(response)
}
