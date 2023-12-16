package utils

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func IsNumber(cursor rune) bool {
	return '0' <= cursor && cursor <= '9'
}

func ConvertToInt(input string) int {
	trimmed := strings.TrimSpace(input)
	nr, err := strconv.Atoi(trimmed)

	if err != nil {
		fmt.Println(input)
		log.Fatal("NOT A NUMBER")
	}
	return nr
}
