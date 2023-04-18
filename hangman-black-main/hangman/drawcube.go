package hangman

import (
	"github.com/nsf/termbox-go"
)

// General function that draws a cube
func Draw_cube(positionX, positionY, verticalLenght, horizontalLenght int, fontColor, backgroundColor termbox.Attribute) {	
	// Initialize variables 
	var borderHorizontal rune = 0x2502
	var borderVertical rune = 0x2500
	var borderTopLeft rune = 0x250C
	var borderTopRight rune = 0x2510
	var borderBotomLeft rune = 0x2514
	var borderBottomRight rune = 0x2518

	// Vertical Line Top : :
	for i := 1; i <= verticalLenght; i++ {
		termbox.SetCell(positionX+i, positionY+0, borderVertical, fontColor, backgroundColor)
	}

	// Vertical Line Bottom :
	for i := 1; i <= verticalLenght; i++ {
		termbox.SetCell(positionX+i, positionY+horizontalLenght+1, borderVertical, fontColor, backgroundColor)
	}

	// Horizontal Line Left :
	for i := 1; i <= horizontalLenght; i++ {
		termbox.SetCell(positionX+0, positionY+i, borderHorizontal, fontColor, backgroundColor)
	}

	// Horizontal Line Right :
	for i := 1; i <= horizontalLenght; i++ {
		termbox.SetCell(positionX+verticalLenght+1, positionY+i, borderHorizontal, fontColor, backgroundColor)
	}

	// Corner Top Left :
	termbox.SetCell(positionX+0, positionY+0, borderTopLeft, fontColor, backgroundColor)
	// Corner Top Right :
	termbox.SetCell(positionX+verticalLenght+1, positionY+0, borderTopRight, fontColor, backgroundColor)
	// Corner Bottom Left :
	termbox.SetCell(positionX+0, positionY+horizontalLenght+1, borderBotomLeft, fontColor, backgroundColor)
	// Corner Bottom Right :
	termbox.SetCell(positionX+verticalLenght+1, positionY+horizontalLenght+1, borderBottomRight, fontColor, backgroundColor)
}
