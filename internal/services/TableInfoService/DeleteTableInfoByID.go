package TableInfoService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteTableInfoByIDRequest struct {
	TableId uuid.UUID `json:"table_id"`
}

func (ti *TableInfoService) DeleteTableInfoByID(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request DeleteTableInfoByIDRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ti.TableInfoRepository.CheckExistByID(ctx.Context(), request.TableId) {
		return utils.SendBadRequest(ctx, &response, "Table Id is not exist")

	}

	err := ti.TableInfoRepository.DeleteOneById(ctx.Context(), request.TableId)

	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	response["message"] = "Deleted"
	return ctx.Status(200).JSON(response)

}
