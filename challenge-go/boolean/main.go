package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven() bool {
	var t string
	mario := os.Args
	for a := 1; a < len(mario); a++ {
		t += mario[a]
	}
	nbr := len(mario)
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	if isEven() == false {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
