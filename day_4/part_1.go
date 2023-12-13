package main

import (
	"fmt"
	"math"
	"utils"
)

func getWinningPoints(input string) {
	lines := utils.SplitInputOnDivider(input, "\n")

	won_total := 0
	for _, line := range lines {
		winning_count := 0
		separated_game := utils.SplitInputOnDivider(line, ":")
		number_fields := utils.SplitInputOnDivider(separated_game[1], "|")
		winning_numbers := utils.SplitInputOnDivider(number_fields[0], " ")
		game_numbers := utils.SplitInputOnDivider(number_fields[1], " ")

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
		if winning_count == 0 {
			continue
		}
		if winning_count > 1 {
			won_total += int(math.Pow(2, float64(winning_count)-1))
		} else {
			won_total += 1
		}
	}

	fmt.Println(won_total)
}
