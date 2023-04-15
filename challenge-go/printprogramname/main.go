package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[0]
	for a, m := range args {
		if a > 1 {
			z01.PrintRune(rune(m))
		}
	}
	z01.PrintRune('\n')
}
