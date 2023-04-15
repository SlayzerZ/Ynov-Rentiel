package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n < 0 {
		if n < -9223372036854775807 {
			PrintNbr(n / 10)
			z01.PrintRune(rune(8) + 48)
			return
		}
		n *= -1
		z01.PrintRune('-')
	}
	if n > 10 {
		PrintNbr(n / 10)
		PrintNbr(n % 10)
	} else {
		z01.PrintRune(rune(n) + 48)
	}
}
