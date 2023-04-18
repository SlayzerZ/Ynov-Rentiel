package hangman

import (
	"github.com/nsf/termbox-go"
)

// Draw the menu for page 1, 2 and 3 (Welcome, Game, Help)
func BarMenu(y int, indexMenu int) {
	// === Menu Bar ===

	// Cube menu Bar :
	Draw_cube(0, y, 55, 1, termbox.ColorWhite, termbox.ColorDefault)

	// Welcome :
	if indexMenu == 0 {
		Draw_text(17, y+1, "Welcome", termbox.ColorLightRed, termbox.ColorDefault)
	} else {
		Draw_text(17, y+1, "Welcome", termbox.ColorWhite, termbox.ColorDefault)
	}
	Draw_text(25, y+1, "|", termbox.ColorWhite, termbox.ColorDefault)

	// Game :
	if indexMenu == 1 {
		Draw_text(27, y+1, "Game", termbox.ColorLightRed, termbox.ColorDefault)
	} else {
		Draw_text(27, y+1, "Game", termbox.ColorWhite, termbox.ColorDefault)
	}
	Draw_text(32, y+1, "|", termbox.ColorWhite, termbox.ColorDefault)

	// Help :
	if indexMenu == 2 {
		Draw_text(34, y+1, "Help", termbox.ColorLightRed, termbox.ColorDefault)
	} else {
		Draw_text(34, y+1, "Help", termbox.ColorWhite, termbox.ColorDefault)
	}
}

// Draw the menu for page 4 and 5 (Win or lose)
func BarMenuResult(y int, indexMenu int) {
	// indexMenu == 0 => Win
	// indexMenu == 1 => Lose

	// Cube menu Bar :
	Draw_cube(0, y, 55, 1, termbox.ColorWhite, termbox.ColorDefault)

	// Lose/Win :
	if indexMenu == 1 {
		Draw_text(25, y+1, "/", termbox.ColorWhite, termbox.ColorDefault)
		Draw_text(27, y+1, "LOSE", termbox.ColorLightRed, termbox.ColorDefault)
		Draw_text(32, y+1, "\\", termbox.ColorWhite, termbox.ColorDefault)
	} else {
		Draw_text(25, y+1, "/", termbox.ColorWhite, termbox.ColorDefault)
		Draw_text(27, y+1, "WIN", termbox.ColorLightGreen, termbox.ColorDefault)
		Draw_text(31, y+1, "\\", termbox.ColorWhite, termbox.ColorDefault)
	}

}
