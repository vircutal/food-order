package TableInfoService

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type DeleteTableInfoByIDRequest struct {
	TableId uuid.UUID `json:"table_id"`
}

func (ti *TableInfoService) DeleteTableInfoByID(ctx *fiber.Ctx) error {
	var request DeleteTableInfoByIDRequest
	response := map[string]string{}

	if err := ctx.BodyParser(&request); err != nil {
		response["message"] = "Server Problem. Try Again..."
		return ctx.JSON(response)
	}

	if !ti.TableInfoRepository.CheckTableIDExist(ctx.Context(), request.TableId) {
		ctx.Status(400)
		response["message"] = "Table Id is not exist"
		return ctx.JSON(response)
	}

	err := ti.TableInfoRepository.DeleteOneById(ctx.Context(), request.TableId)

	if err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
	} else {
		ctx.Status(200)
		response["message"] = "Deleted"
	}
	return ctx.JSON(response)

}
