package piscine

func Capitalize(s string) string {
	tabS := []rune(s)
	IsCapital := true
	for i := 0; i < len(s); i++ {
		if ((tabS[i] >= 'A' && tabS[i] <= 'Z') || (tabS[i] >= 'a' && tabS[i] <= 'z') || (tabS[i] >= '0' && tabS[i] <= '9')) && IsCapital == true {
			if tabS[i] >= 'a' && tabS[i] <= 'z' {
				tabS[i] = 'A' - 'a' + tabS[i]
			}
			IsCapital = false
		} else if tabS[i] >= 'A' && tabS[i] <= 'Z' {
			tabS[i] = 'a' - 'A' + tabS[i]
		} else if !((tabS[i] >= 'A' && tabS[i] <= 'Z') || (tabS[i] >= 'a' && tabS[i] <= 'z') || (tabS[i] >= '0' && tabS[i] <= '9')) {
			IsCapital = true
		}
	}
	return string(tabS)
}
