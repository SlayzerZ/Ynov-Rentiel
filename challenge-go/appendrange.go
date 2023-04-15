package piscine

func AppendRange(min, max int) []int {
	var t []int
	for a := min; a < max; a++ {
		t = append(t, a)
	}
	return t
}
