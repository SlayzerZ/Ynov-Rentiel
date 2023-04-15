package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	trune := []string{}
	for _, b := range args {
		var min rune
		min = rune(b[0])
		index := 0
		for a := 0; a < len(b); a++ {
			if min > rune(b[a]) {
				min = rune(b[a])
				index = a
			}
		}
		trune = append(trune, b[index:])
		for _, c := range trune {
			for _, d := range c {
				z01.PrintRune(d)
			}
		}
	}
}
