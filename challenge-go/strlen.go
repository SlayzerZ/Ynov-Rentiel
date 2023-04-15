package piscine

func StrLen(s string) int {
	for rune(s[1]) > 127 {
		return len(s) - 1
	}
	return len(s)
}
