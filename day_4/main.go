package main

import (
	"fmt"
	"utils"
)

var input = `
Card    1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card  2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card  3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card  4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card  5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card  6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
`

func main() {
	// getWinningPoints(Input)

	lines := utils.SplitInputOnDivider(Input, "\n")

	wonned_card := make([]int, 200)
	for _, line := range lines {
		winning_count := 0
		separated_game := utils.SplitInputOnDivider(line, ":")
		card_nr := utils.FindNumberFromString(separated_game[0])
		number_fields := utils.SplitInputOnDivider(separated_game[1], "|")
		winning_numbers := utils.SplitInputOnDivider(number_fields[0], " ")
		game_numbers := utils.SplitInputOnDivider(number_fields[1], " ")

		wonned_card[card_nr] += 1
		for _, win_nr := range winning_numbers {
			for _, game_nr := range game_numbers {
				if game_nr == "" {
					continue
				}
				if win_nr == game_nr {
					winning_count++
				}
			}

		}

		for i := 0; i < winning_count; i++ {
			wonned_card[card_nr+i+1] += 1 * (wonned_card[card_nr])
		}
	}

	total := 0
	for _, nr := range wonned_card {
		total += nr
	}

	fmt.Println(total)
}
