package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

var input1 string = `
	Game 1: 1 red, 5 blue, 1 green; 16 blue, 3 red; 6 blue, 5 red; 4 red, 7 blue, 1 green
	Game 2: 4 blue; 4 red, 3 blue, 1 green; 4 red, 9 blue, 2 green; 5 blue, 7 green, 4 red
	Game 3: 10 blue; 7 blue, 1 green; 19 blue, 1 green, 9 red
	Game 4: 2 green; 14 blue, 14 red, 4 green; 12 red, 11 green, 13 blue; 5 green, 9 red, 4 blue; 9 red, 7 green, 12 blue; 2 green, 3 blue, 8 red
	Game 5: 3 blue, 4 red; 12 red, 2 green, 15 blue; 1 red, 10 blue, 1 green
	Game 6: 1 blue, 7 red; 3 green, 5 red, 1 blue; 1 green, 7 red; 6 red, 1 blue, 4 green; 1 green, 8 red, 1 blue; 2 green, 4 red, 1 blue
	Game 7: 11 green, 10 blue, 2 red; 1 green, 12 blue, 2 red; 9 green, 14 blue; 1 red, 19 blue, 15 green
	Game 8: 4 green, 2 red, 14 blue; 9 green, 1 red, 15 blue; 2 green, 9 red, 8 blue; 11 green, 7 red, 8 blue; 9 red, 7 green, 6 blue
	Game 9: 4 blue, 1 green, 2 red; 1 blue, 3 red; 1 red, 3 blue, 3 green
	Game 10: 4 red, 3 green, 6 blue; 2 green, 15 blue, 6 red; 3 green, 2 blue; 2 red, 1 green; 11 blue, 7 red, 4 green; 2 blue, 2 red, 4 green
`

func (cfg *GameConfig) isValidGame(game_sets []string) bool {
	for _, set := range game_sets {
		values := strings.Split(set, ",")

		for _, val := range values {
			draw := strings.Split(strings.TrimSpace(val), " ")

			box_count, err := strconv.Atoi(draw[0])
			if err != nil {
				log.Fatalf("Not a number: %s", err)
			}

			box_color := draw[1]

			switch box_color {
			case "red":
				if box_count > cfg.red {
					return false
				}
			case "green":
				if box_count > cfg.green {
					return false
				}
			case "blue":
				if box_count > cfg.blue {
					return false
				}
			}
		}
	}

	return true
}

func partOneRun(all_lines []string) {
	rules := newGameConfig(12, 13, 14)

	total := 0
	for i, line := range all_lines[1:] {
		game := strings.Split(line, ":")
		game_sets := strings.Split(game[1], ";")
		isValid := rules.isValidGame(game_sets)

		if isValid {
			total += i + 1
		}
	}
	fmt.Printf("TOTAL SUMMED WON GAMES: %v \n", total)
}
