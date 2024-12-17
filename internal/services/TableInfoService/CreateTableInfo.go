package TableInfoService

import (
	"food-order/internal/config/constant"
	"food-order/internal/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AddTableInfoRequest struct {
	TableNumber int `json:"table_number"`
}

func (ti *TableInfoService) CreateTableInfo(ctx *fiber.Ctx) error {
	//define request and response
	var request AddTableInfoRequest
	response := map[string]string{}
	//parsing request
	if err := ctx.BodyParser(&request); err != nil {
		response["message"] = err.Error()
		ctx.JSON(response)
		return err
	}
	//check if table exist
	if ti.TableInfoRepository.CheckTableNumberExist(ctx.Context(), request.TableNumber) {
		response["message"] = "Table already exists"
		ctx.Status(400)
		return ctx.JSON(response)
	}

	tmp := models.TableInfo{
		ID:          uuid.New(),
		TableNumber: request.TableNumber,
		Status:      constant.TableIsAvailable,
	}

	// querying TableInfo
	err := ti.TableInfoRepository.AddOne(ctx.Context(), &tmp)
	if err != nil {
		ctx.Status(500)
		response["message"] = err.Error()
	} else {
		ctx.Status(200)
		response["message"] = "table is successfully added"
		response["tableNumber"] = strconv.Itoa(request.TableNumber)
	}
	//
	return ctx.JSON(response)
}
