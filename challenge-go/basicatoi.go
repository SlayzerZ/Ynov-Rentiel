package piscine

func IterativePower2(nb, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	for i := 0; i < power; i++ {
		res *= nb
	}
	return res
}

func BasicAtoi(s string) int {
	nb := 0
	nbs := []int{}
	for a := 0; a < len(s); a++ {
		if s[a] != 0 {
			if s[a] >= 48 && s[a] <= 57 {
				nbs = append(nbs, int(s[a])-48)
			}
		}
		if len(nbs) == 0 {
			return 0
		}
		count := len(nbs) - 1
		f := IterativePower2(10, count)
		nb = nbs[0] * f
		for b := 1; b < len(nbs); b++ {
			f = IterativePower2(10, count-b)
			nb += nbs[b] * f
		}
	}
	return nb
}
