package main

import "github.com/gofiber/fiber/v2"

func (app *Application) OrderRoutes(router fiber.Router) {
	cart := router.Group("/order")
	cart.Post("/checkout", app.ConfirmProductCart)
	cart.Get("/order", app.GetOrder)
}

func (app *Application) ConfirmProductCart(c *fiber.Ctx) error {
	return nil
}

func (app *Application) GetOrder(c *fiber.Ctx) error {
	return nil
}
