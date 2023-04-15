package piscine

func StrRev(s string) string {
	var t string
	for a := len(s) - 1; a >= 0; a-- {
		t += string(s[a])
	}
	return t
}
