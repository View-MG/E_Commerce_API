package main

import (
	"github.com/View-MG/be-project/internal/entity"
	"github.com/gofiber/fiber/v2"
)

func (h *Application) RegisterRoutes(app *fiber.App) {
	app.Post("/orders", h.CreateOrder)
}

func (h *Application) CreateOrder(c *fiber.Ctx) error {
	var order entity.Order
	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	if err := h.service.Order.CreateOrder(order); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(order)
}
