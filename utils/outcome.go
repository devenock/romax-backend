package utils

import (
	"math/rand"
	"time"
)

// Symbols for the slot machine
var symbols = []string{"A", "B", "C", "D", "E"}

// GenerateOutcome simulates a spin by generating random symbols for the reels
func GenerateOutcome() [][]string {
	rand.Seed(time.Now().UnixNano())
	outcome := make([][]string, 3)

	for i := 0; i < 3; i++ {
		outcome[i] = make([]string, 5)
		for j := 0; j < 5; j++ {
			outcome[i][j] = symbols[rand.Intn(len(symbols))]
		}
	}

	return outcome
}
