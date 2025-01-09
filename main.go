package main

import (
	"github.com/devenock/romax-backend/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Initialize the Fiber app
	app := fiber.New()

	// Define routes
	app.Post("/game/spin", handlers.Spin)

	// Start the server
	if err := app.Listen(":2025"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
