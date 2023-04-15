package piscine

func Index(s string, toFind string) int {
	s1 := []rune(s)
	toFind1 := []rune(toFind)
	for i := 0; i <= (len(s1) - len(toFind1)); i++ {
		if toFind == s[i:i+len(toFind1)] {
			return i
		}
	}
	return -1
}
