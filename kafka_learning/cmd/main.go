package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/shutterscripter/kafka_chat/router"
)

func main() {
	app := fiber.New()
	router.SetupRoutes(app)


	

	log.Fatal(app.Listen(":3002"))
}
