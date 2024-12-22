package TableInfoService

import (
	"food-order/internal/config"
	"food-order/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateTableInfoStatusRequest struct {
	TableID uuid.UUID `json:"table_id"`
	Status  string    `json:"status"`
}

func (ti *TableInfoService) UpdateTableInfoStatus(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request UpdateTableInfoStatusRequest
	response := map[string]interface{}{}

	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}
	//**************************************************************
	//**************************************************************

	if !ti.TableInfoRepository.CheckExistByID(ctx.Context(), request.TableID) {
		return utils.SendBadRequest(ctx, &response, "Table doesnt exist")
	}

	targetTable, err := ti.TableInfoRepository.FindOneById(ctx.Context(), request.TableID)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	if !config.TableInfoStatusTransitionRules[targetTable.Status][request.Status] {
		return utils.SendBadRequest(ctx, &response, "Status transition is not valid")
	}

	//update table status
	targetTable.Status = request.Status
	if err = ti.TableInfoRepository.UpdateOne(ctx.Context(), targetTable); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["message"] = "Status changed"
	response["result"] = *targetTable
	return ctx.JSON(response)

}
