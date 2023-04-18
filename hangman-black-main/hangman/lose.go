package hangman

import (
	"github.com/nsf/termbox-go"
)

func Lose() {
	// Function that draws the page when the player has lost :
	Draw_cube(0, 4, 55, 16, termbox.ColorLightRed, termbox.ColorBlack)
	Draw_text(13, 6, "You lose, the word was \"" + Read("hangman/solution/Fword.txt", 1) + "\"", termbox.ColorLightRed, termbox.ColorDefault)
	Draw_text(12, 8, "To restart a game press : \"Enter\"", termbox.ColorWhite, termbox.ColorDefault)
	
	// Draw the ascii art: Winning
	for i := 0; i <= 11; i++ {
		var line = Read("hangman/asciiart/lose.txt", i)
		Draw_text(17, 10+i, line, termbox.ColorLightRed, termbox.ColorDefault)
	}
}
