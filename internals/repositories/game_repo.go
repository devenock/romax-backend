package repositories

import (
	"github.com/devenock/romax-backend/internals/models"
	"github.com/devenock/romax-backend/pkg/database"
)

func CreateGame(game *models.Game) error {
	return database.DB.Create(game).Error
}
