package piscine

func BasicJoin(strs []string) string {
	var r string
	for i := range strs {
		r += strs[i]
	}
	return r
}
