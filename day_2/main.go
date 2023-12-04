package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type GameConfig struct {
	red   int
	green int
	blue  int
}

var input2 string = `
	Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
	Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
	Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`

func newGameConfig(red, green, blue int) GameConfig {
	return GameConfig{
		red:   red,
		green: green,
		blue:  blue,
	}
}

func main() {
	all_lines := strings.Split(Input, "Game")
	// partOneRun(all_lines)
	total := 0
	for _, line := range all_lines[1:] {
		game := strings.Split(line, ":")
		game_sets := strings.Split(game[1], ";")
		total += getLowestSet(game_sets)
	}

	fmt.Println(total)
}

func getLowestSet(game_sets []string) int {
	lowestSet := newGameConfig(0, 0, 0)
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
				if box_count > lowestSet.red {
					lowestSet.red = box_count
				}
			case "green":
				if box_count > lowestSet.green {
					lowestSet.green = box_count
				}
			case "blue":
				if box_count > lowestSet.blue {
					lowestSet.blue = box_count
				}
			}
		}
	}

	return lowestSet.red * lowestSet.green * lowestSet.blue
}
