package piscine

func UltimateDivMod(a *int, b *int) {
	c := *a / *b
	k := *a % *b
	*a = c
	*b = k
}
