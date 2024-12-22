package TableInfoService

import (
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
)

type GetTableInfoByStatusRequest struct {
	Status string `json:"status"`
}

func (ti *TableInfoService) FindAllTableInfoByStatus(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request GetTableInfoByStatusRequest
	response := map[string]interface{}{}

	//parse body
	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	tableInfos, err := ti.TableInfoRepository.GetAllByStatus(ctx.Context(), request.Status)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.JSON(200)
	response["result"] = *tableInfos
	return ctx.JSON(response)
}
