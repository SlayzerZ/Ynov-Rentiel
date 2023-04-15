package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	liste := []int{}
	for n != 0 {
		liste = append(liste, n%10)
		n = n / 10
	}

	for len(liste) > 0 {
		min := liste[0]
		var index int
		for a := 1; a < len(liste); a++ {
			if min > liste[a] {
				min = liste[a]
				index = a
			}
		}
		z01.PrintRune(rune(min) + 48)
		liste = append(liste[:index], liste[index+1:]...)
	}
}
