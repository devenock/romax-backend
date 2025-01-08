package handlers

import (
	"github.com/devenock/romax-backend/internals/services"
	"github.com/devenock/romax-backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func PlayGame(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("playerID"))
	player, err := services.GetPlayerDetails(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Player not found")
	}

	game, err := services.PlayGame(player)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, game)
}
