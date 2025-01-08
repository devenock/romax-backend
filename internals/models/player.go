package models

import "time"

type Player struct {
	ID         uint   `gorm:"primaryKey"`
	Username   string `gorm:"unique;not null"`
	TotalWins  int    `gorm:"default:0"`
	TotalGames int    `gorm:"default:0"`
	CreatedAt  time.Time
}
