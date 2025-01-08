package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"math/rand"
	"time"
)

func main() {
	//	initialize a fiber application
	app := fiber.New()
	app.Use(cors.New())

	//	test endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// game route
	app.Get("/game/result", func(c *fiber.Ctx) error {
		// Simulate clearing symbols
		rand.Seed(time.Now().UnixNano())
		clearCount := rand.Intn(8) // Random number of clears (0 to 7)

		// Determine the result based on clears
		result := "loss"
		freeGames := 0
		if clearCount >= 4 {
			result = "win"
			switch clearCount {
			case 4:
				freeGames = 3
			case 5:
				freeGames = 5
			case 6:
				freeGames = 10
			case 7:
				freeGames = 20
			}
		}

		// Response
		return c.JSON(fiber.Map{
			"status":    "success",
			"result":    result,
			"clears":    clearCount,
			"freeGames": freeGames,
		})
	})

	//	run the server
	app.Listen(":3000")
}
