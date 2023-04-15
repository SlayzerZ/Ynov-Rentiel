package piscine

func BasicJoin(elems []string) string {
	var str string
	var n int
	n = len(elems) - 1

	for a := 0; a <= n; {
		str += elems[a]
		a++
	}
	return str
}
