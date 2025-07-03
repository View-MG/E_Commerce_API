package main

import (
	"fmt"

	dtoReq "github.com/View-MG/be-project/internal/dto/request"
	"github.com/View-MG/be-project/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func (app *Application) AuthRoutes(router *fiber.App) {
	auth := router.Group("/auth")
	auth.Post("/sign-in", app.CreateUser)
	auth.Post("/log-in", app.Login)
}

func (app *Application) CreateUser(c *fiber.Ctx) error {
	var req dtoReq.User

	err := c.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadGateway, fmt.Sprintf("[ERROR][CreateUser] : %v", err))
	}

	res, err := app.Service.User.CreateUser(c, req)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadGateway, fmt.Sprintf("[ERROR][CreateUser] : %v", err))
	}
	return response.SuccessResponse(c, res)
}

func (app *Application) Login(c *fiber.Ctx) error {
	var req dtoReq.LoginRequest

	err := c.BodyParser(&req)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadGateway, fmt.Sprintf("[ERROR][Login] : %v", err))
	}
	res, err := app.Service.User.Login(c, req)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusBadGateway, fmt.Sprintf("[ERROR][Login] : %v", err))
	}
	return response.SuccessResponse(c, res)
}
