package piscine

func AlphaCount(s string) int {
	num := int(0)
	for i := 0; i < len(s); i++ {
		switch {
		case 65 <= s[i] && s[i] <= 90:
			fallthrough
		case 97 <= s[i] && s[i] <= 122:
			num++
		default:
		}
	}
	return num
}
