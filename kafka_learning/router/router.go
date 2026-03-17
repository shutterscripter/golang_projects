package router

import (
	"github.com/gofiber/fiber/v3"
	"github.com/shutterscripter/kafka_chat/handler"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Get("/health", handler.HealthCheck)
	api.Post("/comment", handler.CreateComment)

}
