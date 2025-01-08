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
		rand.Seed(time.Now().UnixNano())
		result := "lose"
		if rand.Intn(2) == 1 { // 50% chance to win
			result = "win"
		}

		// Random win amount (if win occurs)
		winAmount := 0
		if result == "win" {
			winAmount = rand.Intn(100) + 1 // Random win between 1 and 100
		}

		return c.JSON(fiber.Map{
			"status":    "success",
			"result":    result,
			"winAmount": winAmount,
		})
	})

	//	run the server
	app.Listen(":3000")
}
