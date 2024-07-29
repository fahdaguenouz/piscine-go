package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	}
	f := 1
	for i := 1; i < nb; i++ {
		f *= i
	}
	return f

}
