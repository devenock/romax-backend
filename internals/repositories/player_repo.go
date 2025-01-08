package repositories

import (
	"github.com/devenock/romax-backend/internals/models"
	"github.com/devenock/romax-backend/pkg/database"
)

func CreatePlayer(player *models.Player) error {
	return database.DB.Create(player).Error
}

func GetPlayerByID(id uint) (*models.Player, error) {
	var player models.Player
	err := database.DB.First(&player, id).Error
	return &player, err
}

func UpdatePlayer(player *models.Player) error {
	return database.DB.Save(player).Error
}
