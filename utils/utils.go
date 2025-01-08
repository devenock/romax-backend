package utils

import (
	"github.com/gofiber/fiber/v2"
	"math/rand"
	"time"
)

// GenerateRandomNumber generates a random number between min and max.
func GenerateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

// ErrorResponse generates a consistent error response format.
func ErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(fiber.Map{
		"error": message,
	})
}

// SuccessResponse generates a consistent success response format.
func SuccessResponse(c *fiber.Ctx, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": data,
	})
}
