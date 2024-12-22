package TableInfoService

import (
	"food-order/internal/config"
	"food-order/internal/models"
	"food-order/internal/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type AddTableInfoRequest struct {
	TableNumber string `json:"table_number"`
}

func (ti *TableInfoService) CreateTableInfo(ctx *fiber.Ctx) error {
	//initialize instance using in this function
	//**************************************************************
	var request AddTableInfoRequest
	response := map[string]interface{}{}
	//parsing request
	if err := ctx.BodyParser(&request); err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	tableNumber, _ := strconv.Atoi(request.TableNumber)
	//**************************************************************
	//**************************************************************

	//check if table exist
	if ti.TableInfoRepository.CheckTableNumberExist(ctx.Context(), tableNumber) {
		return utils.SendBadRequest(ctx, &response, "Table already exists")
	}

	//init new TableInfo object
	newTableInfo := models.TableInfo{
		ID:          uuid.New(),
		TableNumber: tableNumber,
		Status:      config.TableIsAvailable,
	}

	// Inserting TableInfo
	err := ti.TableInfoRepository.AddOne(ctx.Context(), &newTableInfo)
	if err != nil {
		return utils.SendInternalServerError(ctx, &response, err.Error())
	}

	ctx.Status(200)
	response["message"] = "table is successfully added"
	response["result"] = newTableInfo
	return ctx.JSON(response)
}
