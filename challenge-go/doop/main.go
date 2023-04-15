package main

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	{
		for a := 0; a < len(s); a++ {
			z01.PrintRune(rune(s[a]))
		}
	}
}

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

func IsNumeric3(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 47 || int(letter) >= 58 && int(letter) <= 127 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	var res int
	trune := []rune(s)
	negatif := 1
	for i, c := range trune {
		if c == '-' {
			if i != 0 {
				return 0
			} else {
				negatif = -1
			}

		} else if c == '+' {
			if i != 0 {
				return 0
			}
		} else if !IsNumeric3(string(c)) {
			return 0
		} else {
			if i == len(trune)-1 {
				res += int(c - 48)
			} else {
				res += int(c - 48)
				res *= 10
			}

		}

	}
	return res * negatif
}

func main() {
	args := os.Args
	if len(args) != 4 {
		return
	}
	nb1 := Atoi(args[1])
	nb2 := Atoi(args[3])
	op := args[2]
	if IsNumeric3(string(nb1)) == true || IsNumeric3(string(nb2)) == true {
		return
	}
	switch op {
	case "+":
		PrintNbr(nb1 + nb2)
	case "-":
		PrintNbr(nb1 - nb2)
	case "*":
		PrintNbr(nb1 * nb2)
	case "/":
		if nb2 == 0 {
			PrintStr("No division by 0")
			return
		}
		PrintNbr(nb1 / nb2)
	case "%":
		if nb2 == 0 {
			PrintStr("No modulo by 0")
			return
		}
		PrintNbr(nb1 % nb2)
	}
}
