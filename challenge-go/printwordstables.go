package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		for _, c := range s {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
