package utils

func IsNumber(cursor rune) bool {
	return '0' <= cursor && cursor <= '9'
}
