package models

import "time"

type Player struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Username   string `gorm:"unique;not null" json:"username"`
	TotalWins  int    `gorm:"default:0" json:"total_wins"`
	TotalGames int    `gorm:"default:0" json:"total_games"`
	CreatedAt  time.Time
}
