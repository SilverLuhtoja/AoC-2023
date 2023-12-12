package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"utils"
)

type Table struct {
	matrix [][]rune
}

func newMatrixTable(input string) Table {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	m := make([][]rune, len(lines))

	for i, line := range lines {
		m[i] = []rune(line)
	}

	return Table{matrix: m}
}

func main() {
	table := newMatrixTable(Input)

	gearMap := make(map[string][]int)

	for row := range table.matrix {
		start := 0

		for col := 0; col < len(table.matrix[0]); col++ {
			start = col
			current_number := ""

			for col < len(table.matrix[0]) && utils.IsNumber(table.matrix[row][col]) {
				current_number += string(table.matrix[row][col])
				col += 1
			}

			if current_number == "" {
				continue
			}

			number, err := strconv.Atoi(current_number)
			if err != nil {
				log.Fatal("NOT A NUMBER")
			}

			if table.lookForStar(row, start-1) {
				gearMap[makeKey(row, start-1)] = append(gearMap[makeKey(row, start-1)], number)
				continue
			}

			if table.lookForStar(row, col) {
				gearMap[makeKey(row, col)] = append(gearMap[makeKey(row, col)], number)
				continue
			}

			for c := start - 1; c < col+1; c++ {
				if table.lookForStar(row-1, c) {
					gearMap[makeKey(row-1, c)] = append(gearMap[makeKey(row-1, c)], number)
					break
				}

				if table.lookForStar(row+1, c) {
					gearMap[makeKey(row+1, c)] = append(gearMap[makeKey(row+1, c)], number)
					break
				}
			}
		}
	}

	total := 0

	for _, values := range gearMap {
		if len(values) == 2 {
			total += values[0] * values[1]
		}
	}

	fmt.Println(total)
}

func (t *Table) lookForStar(row, col int) bool {
	if !(0 <= row && row < len(t.matrix) && 0 <= col && col < len(t.matrix[0])) {
		return false
	}

	return t.matrix[row][col] == '*'
}

func makeKey(row, col int) string {
	return fmt.Sprintf(strconv.Itoa(row), "-", strconv.Itoa(col))
}
