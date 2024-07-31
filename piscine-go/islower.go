package piscine

func IsLower(s string) bool {
	str := []rune(s)
	for _, letter := range str {
		if letter < 'a' || letter > 'z' {
			return false
		}
	}
	return true
}
