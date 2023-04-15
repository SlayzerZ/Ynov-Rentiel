package piscine

func IsNumeric(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 47 || int(letter) >= 58 && int(letter) <= 127 {
			return false
		}
	}
	return true
}
