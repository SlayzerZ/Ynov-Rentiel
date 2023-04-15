package main

import (
	"github.com/01-edu/z01"
)

type pointx struct {
	x int
	y int
}

var point pointx

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

func setPoint(x, y int) {
	point.x = 42
	point.y = 21
}

func main() {
	points := &point

	setPoint(points.x, points.y)

	s := ("x = ")
	c := (", y = ")

	for a := 0; a < len(s); a++ {
		z01.PrintRune(rune(s[a]))
	}
	PrintNbr(points.x)
	for b := 0; b < len(c); b++ {
		z01.PrintRune(rune(c[b]))
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}
