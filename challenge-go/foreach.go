package piscine

func ForEach(f func(int), a []int) {
	for n := 0; n < len(a); n++ {
		f(a[n])
	}
}
