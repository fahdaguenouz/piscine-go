package piscine

func IsNumeric(s string) bool {
	str := []rune(s)
	for _, letter := range str {
		if letter > '9' || letter < '0' {
			return false
		}
	}
	return true
}
