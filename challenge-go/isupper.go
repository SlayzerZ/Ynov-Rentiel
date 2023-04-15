package piscine

func IsUpper(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 64 || int(letter) >= 91 && int(letter) <= 127 {
			return false
		}
	}
	return true
}
