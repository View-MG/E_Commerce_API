package main

import (
	"fmt"

	"github.com/View-MG/be-project/internal/adapters"
	"github.com/View-MG/be-project/internal/repository"
	"github.com/View-MG/be-project/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type HttpOrderHandler struct {
	service service.AllOfService
}

func NewHttpOrderHandler(srv service.AllOfService) *HttpOrderHandler {
	return &HttpOrderHandler{service: srv}
}

func main() {
	app := fiber.New()
	dsn := "postgres://myuser:mypassword@localhost:5432/mydatabase?sslmode=disable"
	adapter, err := adapters.NewPostgresAdapter(dsn)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database: %v", err))
	}

	corsMiddleware := cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
		MaxAge:           300,
	})
	app.Use(corsMiddleware)
	repo := repository.NewRepository(adapter.DB)
	srv := service.NewService(repo)

	handler := NewHttpOrderHandler(srv)
	handler.RegisterRoutes(app)

	app.Listen(":8080")

}
