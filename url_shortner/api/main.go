package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/shutterscripter/url_shortner/routes"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolvURL)
	app.Post("/api/v1", routes.ShortnUrl)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	app := fiber.New()
	app.Use(logger.New())
	setupRoutes(app)
	app.Listen(os.Getenv("API_PORT"))

}
