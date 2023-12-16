package utils

import (
	"regexp"
	"strconv"
	"strings"
)

func SplitInputOnDivider(input, mark string) []string {
	return strings.Split(strings.TrimSpace(input), mark)
}

func TrimAllSpace(s string) string {
	re := regexp.MustCompile(`\s+`)
	return re.ReplaceAllString(s, " ")
}

func FindNumberFromString(s string) (i int) {
	// Define a regular expression to match "Card" followed by one or more spaces and a number
	re := regexp.MustCompile(`Card\s+(\d+)`)

	// Find the first match in the string
	match := re.FindStringSubmatch(s)

	if len(match) >= 2 {
		number, err := strconv.Atoi(match[1])
		if err != nil {
			return 0
		}
		i = number
	}
	return
}
