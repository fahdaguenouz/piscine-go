package piscine

func ToLower(s string) string {
	str := []rune(s)
	for i, letter := range str {
		if letter >= 'A' && letter <= 'Z' {
			str[i] = letter + ('a' - 'A')
		}
	}
	return string(str)
}
