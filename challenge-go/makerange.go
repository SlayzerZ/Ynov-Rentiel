package piscine

func MakeRange(min, max int) []int {
	var s []int
	if min >= max {
		return s
	}
	t := make([]int, max-min)
	for a := range t {
		t[a] = min + a
	}
	return t
}
