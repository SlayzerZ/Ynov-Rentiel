package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	for _, b := range args[1:] {
		for _, a := range b {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
