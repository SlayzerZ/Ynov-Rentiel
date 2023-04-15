package piscine

func SplitWhiteSpaces(s string) []string {
	var t []string
	var str string
	spaces := false
	esc := 0
	esct := 0
	no_spaces := ""

	for b := 0; b < len(s); b++ {
		if s[b] == ' ' {
			esct += 1
		}
	}

	for a := 0; a < len(s); a++ {
		if s[a] == ' ' {
			spaces = true
			esc += 1
		} else {
			spaces = false
			str += string(s[a])
		}
		if spaces == true {
			t = append(t, str)
			str = ""
			if esc == esct {
				t = append(t, string(s[a+1:]))
			}
		}
	}
	for i, other := range t {
		if other == no_spaces {
			t = append(t[:i], t[i+1:]...)
		}
	}
	return t
}
