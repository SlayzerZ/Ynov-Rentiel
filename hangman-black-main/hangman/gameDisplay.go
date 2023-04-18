package hangman

import (
	"strconv"
	"github.com/nsf/termbox-go"
)

// Display the game
func GameDisplay(letterClick, allLetterClick string, attempt int) {
	// => Cube : HangMan Game
	Draw_cube(0, 4, 16, 1, termbox.ColorLightYellow, termbox.ColorDefault)
	Draw_text(3, 5, "Hangman game", termbox.ColorLightYellow, termbox.ColorDefault)
	
	// => Attempt :
	Draw_cube(18, 4, 12, 1, termbox.ColorLightRed, termbox.ColorDefault)
	Draw_text(19, 4, " Attempt ", termbox.ColorLightRed, termbox.ColorDefault)

	if attempt >= 6 {
		if attempt >= 9 {
			Draw_text(20, 5, strconv.Itoa(attempt), termbox.ColorLightRed, termbox.ColorDefault)
		} else {
			Draw_text(20, 5, strconv.Itoa(attempt), termbox.ColorLightYellow, termbox.ColorDefault)
		}
	} else {
		Draw_text(20, 5, strconv.Itoa(attempt), termbox.ColorLightGreen, termbox.ColorDefault)
	}

	// => Word or Letter :
	Draw_cube(0, 7, 30, 1, termbox.ColorDarkGray, termbox.ColorDefault)
	Draw_text(1, 7, " Letter (or Word) ", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 8, letterClick, termbox.ColorWhite, termbox.ColorDefault)

	// => Word try :
	Draw_cube(0, 10, 30, 1, termbox.ColorDarkGray, termbox.ColorDefault)
	Draw_text(1, 10, " Word ", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 11, Read("hangman/solution/Dword.txt", 1), termbox.ColorWhite, termbox.ColorDefault)
	
	// => Used Letter :
	Draw_cube(0, 13, 30, 3, termbox.ColorDarkGray, termbox.ColorDefault)
	Draw_text(1, 13, " Used Letter ", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 14, allLetterClick, termbox.ColorWhite, termbox.ColorDefault)
	
	// => HangMan Draw :
	Draw_cube(32, 4, 29, 12, termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(33, 4, " Hangman ", termbox.ColorWhite, termbox.ColorDefault)
}
