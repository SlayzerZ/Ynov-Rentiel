package piscine

func IsSorted(f func(a, b int) int, c []int) bool {
	for n := 0; n < len(c)-1; n++ {
		if c[0] < c[1] {
			if f(c[n+1], c[n]) < 0 {
				return false
			}
		} else {
			if f(c[n], c[n+1]) < 0 {
				return false
			}
		}
	}
	return true
}
