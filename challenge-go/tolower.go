package piscine

func ToLower(s string) string {
	var n int
	n = len(s) - 1
	t1 := []rune(s)
	for a := 0; a <= n; a++ {
		if t1[a] >= 65 && t1[a] <= 90 {
			t1[a] += 32
		}
	}
	return string(t1)
}
