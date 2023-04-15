package piscine

func IterativeFactorial(nb int) int {
	n := 1
	if nb < 0 || nb > 20 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	for a := nb; a >= 0; a-- {
		if a != 0 {
			n *= a
		}
	}
	return n
}
