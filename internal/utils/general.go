package utils

import "github.com/gofiber/fiber/v2"

func SendBadRequest(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	(*response)["message"] = message
	return ctx.Status(400).JSON(*response)
}

func SendInternalServerError(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	(*response)["message"] = message
	return ctx.Status(500).JSON(*response)
}

func SendStatusOK(ctx *fiber.Ctx, response *map[string]interface{}, message string) error {
	(*response)["message"] = message
	return ctx.Status(200).JSON(*response)
}
