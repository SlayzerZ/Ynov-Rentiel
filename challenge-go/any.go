package piscine

func Any(f func(string) bool, a []string) bool {
	B := true
	for n := 0; n < len(a); n++ {
		B = f(a[n])
		if B == true {
			return B
		}
	}
	return B
}
