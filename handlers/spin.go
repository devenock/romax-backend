package handlers

import (
	"github.com/devenock/romax-backend/utils"
	"github.com/gofiber/fiber/v2"
)

// Spin handles the spin request and returns the game outcome
func Spin(c *fiber.Ctx) error {
	// Generate random outcome
	outcome := utils.GenerateOutcome()

	// Check paylines and calculate winnings
	matchedPaylines, winnings := utils.CheckPaylines(outcome)

	// Build response
	result := map[string]interface{}{
		"outcome":  outcome,
		"paylines": matchedPaylines,
		"winnings": winnings,
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"result": result,
	})
}
