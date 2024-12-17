package TableInfoService

import (
	"food-order/internal/models"

	"github.com/gofiber/fiber/v2"
)

type GetTableInfoByStatusRequest struct {
	Status string `json:"status"`
}

func (ti *TableInfoService) GetTableInfoByStatus(ctx *fiber.Ctx) error {
	var req GetTableInfoByStatusRequest

	if err := ctx.BodyParser(&req); err != nil {
		return err
	}
	tableInfos, err := ti.TableInfoRepository.GetAllByStatus(ctx.Context(), req.Status)
	if err != nil {
		return err
	}
	data := map[string][]models.TableInfo{
		"message": *tableInfos,
	}

	return ctx.JSON(data)
}
