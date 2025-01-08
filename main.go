package main

import "github.com/gofiber/fiber/v2"

func main() {
	//	initialize a fiber application
	app := fiber.New()

	//	test endpoint
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	//	run the server
	app.Listen(":3000")
}
