package utils

// Define winning paylines
var paylines = [][]int{
	{0, 1, 2, 3, 4},
	{5, 6, 7, 8, 9},
	{10, 11, 12, 13, 14},
}

// Define payouts for matching paylines
var payouts = map[int]int{
	0: 10,
	1: 20,
	2: 30,
}

// CheckPaylines checks the outcome against the paylines and calculates winnings
func CheckPaylines(outcome [][]string) ([]int, int) {
	matchedPaylines := []int{}
	winnings := 0

	// Flatten the outcome for easy payline checking
	flattened := flattenOutcome(outcome)

	for index, payline := range paylines {
		if isWinningPayline(flattened, payline) {
			matchedPaylines = append(matchedPaylines, index)
			winnings += payouts[index]
		}
	}

	return matchedPaylines, winnings
}

// flattenOutcome flattens the 2D outcome into a 1D slice
func flattenOutcome(outcome [][]string) []string {
	flattened := []string{}
	for _, row := range outcome {
		flattened = append(flattened, row...)
	}
	return flattened
}

// isWinningPayline checks if a payline matches the same symbol
func isWinningPayline(flattened []string, payline []int) bool {
	firstSymbol := flattened[payline[0]]
	for _, index := range payline {
		if flattened[index] != firstSymbol {
			return false
		}
	}
	return true
}
