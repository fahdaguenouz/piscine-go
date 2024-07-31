package piscine

func IsPrintable(s string) bool {
	str := []rune(s)
	for _, letter := range str {
		if letter < 32 || letter > 126 {
			return false
		}
	}
	return true
}
