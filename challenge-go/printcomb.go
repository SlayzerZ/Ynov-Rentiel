package piscine

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for indexThree := 48; indexThree <= 57; indexThree++ {
		for indexTwo := 49; indexTwo <= 56; indexTwo++ {
			for indexOne := 48; indexOne <= 57; indexOne++ {
				if indexThree < indexTwo && indexTwo < indexOne {
					z01.PrintRune(rune(indexThree))
					z01.PrintRune(rune(indexTwo))
					z01.PrintRune(rune(indexOne))
					if indexThree == 55 && indexTwo == 56 && indexOne == 57 {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
