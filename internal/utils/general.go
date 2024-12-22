package utils

import "github.com/gofiber/fiber/v2"

func SendBadRequest(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	ctx.Status(400)
	(*response)["message"] = message
	return ctx.JSON(*response)
}

func SendInternalServerError(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	ctx.Status(500)
	(*response)["message"] = message
	return ctx.JSON(*response)
}

func SendStatusOK(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	ctx.Status(200)
	(*response)["message"] = message
	return ctx.JSON(*response)
}
