package piscine

func ToUpper(s string) string {
	str := []rune(s)
	for i, letter := range str {
		if letter >= 'a' && letter <= 'z' {
			str[i] = letter - ('a' - 'A')
		}
	}
	return string(str)
}
