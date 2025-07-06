package main

import "github.com/gofiber/fiber/v2"

func (app *Application) ProductRoutes(router fiber.Router) {
	product := router.Group("/product")
	product.Get("/", app.GetProduct)
	product.Post("/create", app.CreateProduct)
	product.Patch("/:id", app.UpdateProduct)
}

func (app *Application) GetProduct(c *fiber.Ctx) error {
	return nil
}

func (app *Application) CreateProduct(c *fiber.Ctx) error {
	return nil
}

func (app *Application) UpdateProduct(c *fiber.Ctx) error {
	return nil
}
