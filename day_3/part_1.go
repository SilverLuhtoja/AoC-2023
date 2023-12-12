package main

// import (
// 	"fmt"
// 	"log"
// 	"strconv"
// )

// func main() {
// 	table := newMatrixTable(Input)
// 	symboled_numbers := []string{}
// 	for row := range table.matrix {
// 		start := 0
// 		for col := 0; col < len(table.matrix[0]); col++ {
// 			start = col
// 			current_number := ""
// 			for col < len(table.matrix[0]) && (table.matrix[row][col]) {
// 				current_number += string(table.matrix[row][col])
// 				col += 1
// 			}
// 			if current_number == "" {
// 				continue
// 			}
// 			if table.lookSymbols(row, start-1) || table.lookSymbols(row, col) {
// 				symboled_numbers = append(symboled_numbers, current_number)
// 				continue
// 			}
// 			for c := start - 1; c < col+1; c++ {
// 				if table.lookSymbols(row-1, c) || table.lookSymbols(row+1, c) {
// 					symboled_numbers = append(symboled_numbers, current_number)
// 					break
// 				}
// 			}
// 		}
// 	}
// 	total := 0
// 	for _, nr := range symboled_numbers {
// 		number, err := strconv.Atoi(nr)
// 		if err != nil {
// 			log.Fatal("NOT A NUMBER")
// 		}
// 		total += number
// 	}
// 	fmt.Println(total)
// }
// func (t *Table) lookSymbols(row, col int) bool {
// 	if !(0 <= row && row < len(t.matrix) && 0 <= col && col < len(t.matrix[0])) {
// 		return false
// 	}
// 	return !isNumber(t.matrix[row][col]) && t.matrix[row][col] != '.'
// }
