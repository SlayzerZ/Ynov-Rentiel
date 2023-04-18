package hangman

import (
	"github.com/nsf/termbox-go"
)

func Help() {
	// Help Cube :
	Draw_cube(0, 4, 55, 8, termbox.ColorLightBlue, termbox.ColorDefault)
	Draw_text(4, 4, " HELP ", termbox.ColorLightBlue, termbox.ColorDefault)
	
	// Content Help :
	Draw_text(2, 5, "\"ESC\" : Quit Game", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 6, "\"Right Arrow\" : Switch menu", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 7, "\"Left Arrow\" : Switch menu", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 9, "Game :", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 10, "\"Enter\" : Confirm letter", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 11, "\"BackSpace\" : Delete letter", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(2, 12, "\"Space\" : Delete letter", termbox.ColorWhite, termbox.ColorDefault)
}
