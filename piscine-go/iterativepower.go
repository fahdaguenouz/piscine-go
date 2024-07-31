package piscine

func IterativePower(nb int, power int) int {
	r := 1
	if power < 0 {
		r = 0
	} else {
		for i := 0; i < power; i++ {
			r *= nb
		}
	}
	return r
}
