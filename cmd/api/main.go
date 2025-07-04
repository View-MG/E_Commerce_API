package main

import (
	"fmt"

	"github.com/View-MG/be-project/config"
	"github.com/View-MG/be-project/internal/adapters"
	"github.com/View-MG/be-project/internal/repository"
	"github.com/View-MG/be-project/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type Application struct {
	Service   service.AllOfService
	JWTSecret []byte
}

func main() {
	appConfig := config.GetAppConfig()

	postgresInfo := appConfig.Postgres
	app := fiber.New()
	adapter, err := adapters.NewPostgresAdapter(postgresInfo.Username, postgresInfo.Password, postgresInfo.Host, postgresInfo.Port, postgresInfo.DBName)
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

	handler := &Application{
		Service:   srv,
		JWTSecret: []byte(appConfig.JWTSecret),
	}
	handler.AuthRoutes(app)
	api := app.Group("/api", JWTMiddleware(handler.JWTSecret))

	handler.RegisterRoutes(api)
	app.Listen(":" + appConfig.ServerPort)

}
