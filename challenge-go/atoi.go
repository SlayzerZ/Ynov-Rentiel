package piscine

func IsNumeric2(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 47 || int(letter) >= 58 && int(letter) <= 127 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	var res int
	trune := []rune(s)
	negatif := 1
	for i, c := range trune {
		if c == '-' {
			if i != 0 {
				return 0
			} else {
				negatif = -1
			}

		} else if c == '+' {
			if i != 0 {
				return 0
			}
		} else if !IsNumeric2(string(c)) {
			return 0
		} else {
			if i == len(trune)-1 {
				res += int(c - 48)
			} else {
				res += int(c - 48)
				res *= 10
			}

		}

	}
	return res * negatif
}
