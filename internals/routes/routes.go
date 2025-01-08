package routes

import (
	"github.com/devenock/romax-backend/internals/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Post("/player/create", handlers.CreatePlayer)
	app.Get("/player/:id", handlers.GetPlayerDetails)
	app.Get("/game/result/:playerID", handlers.PlayGame)
}
