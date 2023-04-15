package piscine

func Join(strs []string, sep string) string {
	var str string
	var n int
	n = len(strs) - 1

	for a := 0; a <= n; {
		if strs[a] == strs[n] {
			str += strs[a]
			a++
		} else {
			str += strs[a] + sep
			a++
		}
	}
	return str
}
