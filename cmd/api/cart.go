package main

import "github.com/gofiber/fiber/v2"

func (app *Application) CartRoutes(router fiber.Router) {
	cart := router.Group("/cart")
	cart.Get("/", app.GetProductCart)
	cart.Post("/add", app.AddProductCart)
	cart.Patch("/add/{item_id}", app.UpdateProductCart)
	cart.Delete("/add/{item_id}", app.DeleteProductCart)
}

func (app *Application) GetProductCart(c *fiber.Ctx) error {
	return nil
}

func (app *Application) AddProductCart(c *fiber.Ctx) error {
	return nil
}

func (app *Application) UpdateProductCart(c *fiber.Ctx) error {
	return nil
}

func (app *Application) DeleteProductCart(c *fiber.Ctx) error {
	return nil
}
