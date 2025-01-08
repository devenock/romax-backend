package models

import "time"

type Game struct {
	ID        uint   `gorm:"primaryKey"`
	PlayerID  uint   `gorm:"not null"`
	Result    string `gorm:"not null"`
	Clears    int    `gorm:"not null"`
	FreeGames int    `gorm:"not null"`
	CreatedAt time.Time
}
