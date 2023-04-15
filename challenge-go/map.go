package piscine

func Map(f func(int) bool, a []int) []bool {
	B := make([]bool, len(a))
	for n := 0; n < len(a); n++ {
		B[n] = f(a[n])
	}
	return B
}
