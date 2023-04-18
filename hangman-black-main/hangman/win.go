package hangman

import (
	"github.com/nsf/termbox-go"
)

func Win() {
	// Function that draws the page when the player has win the game :
	Draw_cube(0, 4, 55, 16, termbox.ColorLightGreen, termbox.ColorBlack)
	Draw_text(13, 6, "You Win, the word was \"" + Read("hangman/solution/Fword.txt", 1) + "\"", termbox.ColorLightGreen, termbox.ColorDefault)
	Draw_text(12, 8, "To restart a game press : \"Enter\"", termbox.ColorWhite, termbox.ColorDefault)

	// Loop to draw an ascii art : Lost
	for i := 0; i <= 11; i++ {
		var line = Read("hangman/asciiart/win.txt", i)
		Draw_text(12, 8+i, line, termbox.ColorLightGreen, termbox.ColorDefault)
	}
}
