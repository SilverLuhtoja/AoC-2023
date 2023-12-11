package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

// var input1 = `
// 	467..114..
// 	...*......
// 	..35..633.
// 	......#...
// 	617*......
// 	.....+.58.
// 	..592.....
// 	......755.
// 	...$.*....
// 	.664.598..
// `

var input1 = `
	.467.114..
	....*.....
	..35..633.
	...831&267
`

// var input1 = `
// ...766.......821.547.....577......................................387.....................56..........446.793..........292..................
// ...........................%...../.....981..........627..../..........-.....623......610..-..............*..................16......891.....
// ...$...........716..&336.......470.325.................*.84........$..34....*.....+.....#.....*76....#.........303.433........-........&....
// `

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

// func (t *Table) printMatrix(matrix [][]rune) {
// 	for _, row := range matrix {
// 		fmt.Println(string(row))
// 	}
// }

// func (t *Table) isOutOfBounds(row, col int) bool {
// 	if row <= 0 || row > len(t.matrix)-1 || col <= 0 || col > len(t.matrix)-1 {
// 		return true
// 	}
// 	return false
// }

func main() {
	// table := newMatrixTable(Input)
	table := newMatrixTable(input1)

	checkTo := 0
	current_number := ""
	checkFrom := 0
	allNumbersNextToSymbols := []string{}
	for row := range table.matrix {
		for col := range table.matrix[row] {
			cursor := table.matrix[row][col]
			if isNumber(cursor) {
				if current_number == "" {
					if col == 0 {
						checkFrom = col
					} else {
						checkFrom = col - 1
					}
				}
				current_number += string(cursor)

			} else {
				checkTo = col
				if current_number != "" {

					if table.numberNextToSymbol(row, checkFrom, checkTo) {
						allNumbersNextToSymbols = append(allNumbersNextToSymbols, current_number)
					} else {
						fmt.Println(current_number)
						fmt.Println(checkFrom, checkTo)
					}
				}

				current_number = ""
				checkFrom = 0
				checkTo = 0

			}
		}
	}

	total := 0

	for _, nr := range allNumbersNextToSymbols {
		number, err := strconv.Atoi(nr)
		if err != nil {
			log.Fatal("NOT A NUMBER")
		}
		total += number
	}

	// fmt.Println(total)
	fmt.Println(allNumbersNextToSymbols)
}

func isNumber(cursor rune) bool {
	if cursor >= 48 && cursor <= 57 {
		return true
	}
	return false
}

func isSymbol(cursor rune) bool {
	ok := !isNumber(cursor) && string(cursor) != "."
	// if ok {
	// 	fmt.Print(string("GOT SYMBOL :"))
	// 	fmt.Println(string(cursor))
	// }
	return ok
}

func (t *Table) numberNextToSymbol(row, checkFrom, checkTo int) bool {
	if isSymbol(t.matrix[row][checkFrom]) || isSymbol(t.matrix[row][checkTo]) {
		return true
	}

	searcableRow := row - 1
	if searcableRow > 0 {
		// fmt.Printf("UPPER row : %v, from: %v , to: %v \n", searcableRow, checkFrom, checkTo)
		for i := checkFrom; i <= checkTo; i++ {
			if isSymbol(t.matrix[searcableRow][i]) {
				return true
			}
		}
	}

	searcableRow = row + 1
	if searcableRow < len(t.matrix) {
		// fmt.Printf("LOWER row : %v, from: %v , to: %v \n", searcableRow, checkFrom, checkTo)
		for i := checkFrom; i <= checkTo; i++ {
			if isSymbol(t.matrix[searcableRow][i]) {
				return true
			}
		}
	}

	return false
}
