package utils

import "strings"

func SplitInputOnDivider(input, mark string) []string {
	return strings.Split(strings.TrimSpace(input), mark)
}

func TrimAllSpace(s string) string {
	return strings.Join(strings.Fields(s), " ")
}
