package services

import (
	"github.com/devenock/romax-backend/internals/models"
	"github.com/devenock/romax-backend/internals/repositories"
	"github.com/devenock/romax-backend/utils"
)

func PlayGame(player *models.Player) (*models.Game, error) {
	clearCount := utils.GenerateRandomNumber(0, 7)

	result := "loss"
	freeGames := 0
	if clearCount >= 4 {
		result = "win"
		freeGames = map[int]int{4: 3, 5: 5, 6: 10, 7: 20}[clearCount]
		player.TotalWins++
	}

	player.TotalGames++
	if err := repositories.UpdatePlayer(player); err != nil {
		return nil, err
	}

	game := &models.Game{
		PlayerID:  player.ID,
		Result:    result,
		Clears:    clearCount,
		FreeGames: freeGames,
	}
	err := repositories.CreateGame(game)
	return game, err
}
