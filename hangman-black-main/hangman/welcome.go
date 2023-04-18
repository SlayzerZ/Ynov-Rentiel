package hangman

import (
	"github.com/nsf/termbox-go"
)

func Welcome() {
	// Function that draws the page when the player has arrived on the game or when he goes to the Welcome page :

	// Draw the cube
	Draw_cube(0, 4, 64, 22, termbox.ColorLightGreen, termbox.ColorDefault)

	// Draw the house of welcome
	for i := 0; i <= 19; i++ {
		var line = Read("hangman/asciiart/house.txt", i)
		Draw_text(3, 5+i, line, termbox.ColorRed, termbox.ColorDefault)
	}

	// Draw the information messages 
	Draw_text(4, 4, " WELCOME ", termbox.ColorLightGreen, termbox.ColorDefault)
	Draw_text(2, 5, "Welcome to the Hangman!", termbox.ColorLightYellow, termbox.ColorDefault)
	Draw_text(2, 6, "Good luck to you ^^", termbox.ColorDarkGray, termbox.ColorDefault)

	Draw_text(55, 5, "ATTENTION", termbox.ColorLightRed, termbox.ColorDefault)
	Draw_text(47, 6, "Press the \"Enter\"", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(49, 7, "to confirm your", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(56, 8, "choice !", termbox.ColorWhite, termbox.ColorDefault)

	// Credit
	Draw_text(2, 26, "by Thomas & Ibrahim", termbox.ColorDarkGray, termbox.ColorDefault)
}
