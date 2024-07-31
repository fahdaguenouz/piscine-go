package piscine

func IsUpper(s string) bool {
	str := []rune(s)
	for _, letter := range str {
		if letter < 'A' || letter > 'Z' {
			return false
		}
	}
	return true
}
