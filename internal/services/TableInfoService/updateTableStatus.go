package TableInfoService

import (
	"food-order/internal/config/constant"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UpdateTableInfoRequest struct {
	TableID uuid.UUID `json:"table_id"`
	Status  string    `json:"status"`
}

// Not done Yet
func (ti *TableInfoService) UpdateTableInfo(ctx *fiber.Ctx) error {
	var request UpdateTableInfoRequest
	response := map[string]string{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	targetTable, err := ti.TableInfoRepository.FindOneById(ctx.Context(), request.TableID)

	if err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
		return ctx.JSON(response)
	}

	if constant.TableInfoStatusTransitionRules[targetTable.Status][request.Status] {
		targetTable.Status = request.Status
		if err = ti.TableInfoRepository.UpdateOne(ctx.Context(), targetTable); err != nil {
			ctx.Status(500)
			response["message"] = "Server problem. Try Again..."
		} else {
			ctx.Status(200)
			response["message"] = "Status changed"
		}

	} else {
		ctx.Status(400)
		response["message"] = "Status transition is not valid"
	}

	return ctx.JSON(response)

}
