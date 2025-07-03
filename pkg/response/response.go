package main

import (
	"github.com/gofiber/fiber/v2"
)

// Response represents the standard API response structure.
// @swagger:response
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Result  interface{} `json:"result,omitempty"`
}

func SuccessResponse(c *fiber.Ctx, message string, result interface{}) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Message: message,
		Result:  result,
	})
}

func CreatedResponse(c *fiber.Ctx, message string, result interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(Response{
		Success: true,
		Message: message,
		Result:  result,
	})
}

func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(Response{
		Success: false,
		Message: message,
	})
}

func UnprocessableEntity(c *fiber.Ctx, message string, result interface{}) error {
	return c.Status(fiber.StatusUnprocessableEntity).JSON(Response{
		Success: false,
		Message: message,
		Result:  result,
	})
}
