package piscine

func Join(strs []string, sep string) string {
	var r string
	var l int

	for j := range strs {
		j++
		l++
	}

	for i := range strs {
		r += strs[i]
		if i != l-1 {
			r += sep
		}
	}
	return r
}
