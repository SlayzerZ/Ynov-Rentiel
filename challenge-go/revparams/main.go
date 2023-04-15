package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	t := []string{}
	for c := 1; c < len(args); c++ {
		t = append(t, args[len(args)-c])
	}

	for _, b := range t {
		for _, a := range b {
			z01.PrintRune(a)
		}
		z01.PrintRune('\n')
	}
}
