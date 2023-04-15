package piscine

func CountIf(f func(string) bool, tab []string) int {
	B := 0
	for n := 0; n < len(tab); n++ {
		if f(tab[n]) == true {
			B += 1
		}
	}
	return B
}
