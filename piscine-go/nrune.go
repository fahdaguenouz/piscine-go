package piscine

func NRune(s string, n int) rune {
	st := []rune(s)
	if n <= 0 || n > len(st) {
		return 0
	}
	return st[n-1]
}
