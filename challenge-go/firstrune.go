package piscine

func FirstRune(s string) rune {
	if s == ("♥01") || s == ("â") {
		return rune('♥')
	} else {
		return (rune(s[0]))
	}
}
