package piscine

func IsPrintable(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 31 {
			return false
		}
	}
	return true
}
