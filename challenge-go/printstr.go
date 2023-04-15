package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	{
		var n int
		n = len(s) - 1
		for a := 0; a <= n; a++ {
			z01.PrintRune(rune(s[a]))
		}
	}
}
