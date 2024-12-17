package TableInfoService

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Request struct {
	TableID uuid.UUID `json:"table_id"`
	Status  string    `json:"status"`
}

// Not done Yet
func (ti *TableInfoService) UpdateTableInfo(ctx *fiber.Ctx) error {
	var request Request
	response := map[string]string{}

	if err := ctx.BodyParser(&request); err != nil {
		return err
	}

	targetTable, err := ti.TableInfoRepository.FindOneById(ctx.Context(), request.TableID)

	if err != nil {
		fmt.Println(err)
		response["message"] = err.Error()
		ctx.JSON(response)
	} else {

		response["message"] = "Success"
		response["tableStatus"] = targetTable.Status
		response["tableNumber"] = fmt.Sprintf("%d", targetTable.TableNumber)
		ctx.JSON(response)
	}
	return err

}
