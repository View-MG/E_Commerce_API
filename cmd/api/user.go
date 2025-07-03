package main

import (
	dtoReq "github.com/View-MG/be-project/internal/dto/request"
	"github.com/gofiber/fiber/v2"
)

func (app *Application) AuthRoutes(router *fiber.App) {
	auth := router.Group("/auth")
	auth.Post("/sign-in", app.CreateUser)
}

func (app *Application) CreateUser(c *fiber.Ctx) error {
	var req dtoReq.User

	err := c.BodyParser(&req)
	if err != nil {
		
	}
	return c.SendStatus(fiber.StatusCreated)
}
