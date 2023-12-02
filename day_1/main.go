package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

var spelled_numbers = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
var numerics = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

func main() {
	var total int = 0

	for _, line := range strings.Split(Input, "\n") {
		gathered_numbers := []int{}
		current_word := ""

		for _, char := range line {
			if unicode.IsNumber(rune(char)) {
				converted_words := checkForNumbers(current_word)
				gathered_numbers = append(gathered_numbers, converted_words...)
				current_word = ""

				nr, err := strconv.Atoi(string(char))
				if err != nil {
					log.Fatalf("not a number : %v", err)
				}

				gathered_numbers = append(gathered_numbers, nr)
			} else {
				current_word += string(char)
			}
		}

		converted_words := checkForNumbers(current_word)
		gathered_numbers = append(gathered_numbers, converted_words...)

		if len(gathered_numbers) == 0 {
			continue
		}
		total += gathered_numbers[0]*10 + gathered_numbers[len(gathered_numbers)-1]
	}

	fmt.Println(total)
}

func checkForNumbers(word string) (list []int) {
	found_combinations := []int{}
	substring := ""

	for i := 0; i < len(word); i++ {
		substring += string(word[i])
		for _, val := range spelled_numbers {
			if strings.Contains(substring, val) {
				idx := getIndex(val)
				found_combinations = append(found_combinations, numerics[idx])

				substring = substring[len(substring)-2:]
			}
		}

	}
	return found_combinations
}

func getIndex(value string) (index int) {
	for i, val := range spelled_numbers {
		if val == value {
			index = i
		}
	}
	return index
}
