package handlers

import (
	"github.com/devenock/romax-backend/internals/services"
	"github.com/devenock/romax-backend/utils"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func CreatePlayer(c *fiber.Ctx) error {
	type Request struct {
		Username string `json:"username"`
	}

	req := new(Request)
	if err := c.BodyParser(req); err != nil {
		return utils.ErrorResponse(c, fiber.StatusBadRequest, "Invalid input")
	}

	player, err := services.CreatePlayer(req.Username)
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusInternalServerError, err.Error())
	}

	return utils.SuccessResponse(c, player)
}

func GetPlayerDetails(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	player, err := services.GetPlayerDetails(uint(id))
	if err != nil {
		return utils.ErrorResponse(c, fiber.StatusNotFound, "Player not found")
	}

	return utils.SuccessResponse(c, player)
}
