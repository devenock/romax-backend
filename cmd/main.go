package main

import (
	"github.com/devenock/romax-backend/config"
	"github.com/devenock/romax-backend/internals/routes"
	"github.com/devenock/romax-backend/pkg/database"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// Load configuration
	config.LoadConfig()

	// Initialize the database
	database.InitDB()

	// Initialize the app
	app := fiber.New()

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	log.Println("Server running on http://localhost:3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
