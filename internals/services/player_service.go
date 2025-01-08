package services

import (
	"errors"
	"github.com/devenock/romax-backend/internals/models"
	"github.com/devenock/romax-backend/internals/repositories"
)

func CreatePlayer(username string) (*models.Player, error) {
	player := &models.Player{Username: username}
	err := repositories.CreatePlayer(player)
	return player, err
}

func GetPlayerDetails(id uint) (*models.Player, error) {
	player, err := repositories.GetPlayerByID(id)
	if err != nil {
		return nil, errors.New("player not found")
	}
	return player, nil
}
