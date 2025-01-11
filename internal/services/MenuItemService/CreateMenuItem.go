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
	MenuID              uuid.UUID             `form:"menu_id"`
	MenuItemName        string                `form:"menu_item_name"`
	MenuItemPrice       float64               `form:"menu_item_price"`
	MenuItemDescription *string               `form:"menu_item_description"`
	MenuItemImageURL    *multipart.FileHeader `form:"menu_item_image_url"`
}

func (mi *MenuItemService) CreateMenuItem(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request CreateMenuItemRequest
	response := map[string]interface{}{}
	var key *string

	menuRepository := repositories.GetMenuRepository()
	StorageClient := utils.GetStorageClient()

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	menu_item_image_file_header, _ := ctx.FormFile("menu_item_image_url")
	menuItemID := uuid.New()

	if !menuRepository.CheckExistByID(ctx.Context(), request.MenuID) {
		return utils.SendBadRequest(ctx, &response, "MenuID is not exist")
	}
	//**************************************************************
	//**************************************************************

	if menu_item_image_file_header != nil {
		key = utils.GenerateKeyFromFilename(menu_item_image_file_header.Filename, menuItemID.String())
		menu_item_image, _ := menu_item_image_file_header.Open()
		if err := StorageClient.S3PutObject(ctx.Context(), "my-bucket", *key, menu_item_image); err != nil {
			return utils.SendInternalServerError(ctx, &response, err.Error())
		}
	} else {
		key = nil
	}

	targetMenuItem := &models.MenuItem{
		ID:                  menuItemID,
		MenuID:              request.MenuID,
		MenuItemName:        request.MenuItemName,
		MenuItemPrice:       request.MenuItemPrice,
		MenuItemDescription: request.MenuItemDescription,
		MenuItemImageKey:    key,
	}

	if err := mi.MenuItemRepository.AddOne(ctx.Context(), targetMenuItem); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["result"] = *targetMenuItem
	return ctx.Status(201).JSON(response)
}
