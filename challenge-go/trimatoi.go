package piscine

func TrimAtoi(s string) int {
	var long string
	puissance := 1
	var resultat int
	neg := false
	for i := 0; i < len(s); i++ {
		if long == "" {
			if s[i] == 45 {
				neg = true
			}
		}
		if s[i] >= 48 && s[i] <= 57 {
			long += string(s[i])
		}
	}
	for j := 0; j < len(long); j++ {
		puissance = puissance * 10
	}
	puissance /= 10
	for k := 0; k < len(long); k++ {
		resultat += int(long[k]-48) * puissance
		puissance /= 10
	}
	if neg {
		return resultat * -1
	} else {
		return resultat
	}
}
