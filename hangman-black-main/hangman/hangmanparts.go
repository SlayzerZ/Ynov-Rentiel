package hangman

import (
	"github.com/nsf/termbox-go"
)

// ==> All the functions to draw the different parts of the hangman 
func baseHangMan(positionX, positionY int) {
	Draw_text(positionX+1, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+2, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+3, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+4, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+5, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+6, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+7, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+8, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+9, positionY+7, "=", termbox.ColorWhite, termbox.ColorDefault)
}

func beam(positionX, positionY int) {
	Draw_text(positionX+7, positionY+6, "|", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+7, positionY+5, "|", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+7, positionY+4, "|", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+7, positionY+3, "|", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+7, positionY+2, "|", termbox.ColorWhite, termbox.ColorDefault)
}

func HighBeam(positionX, positionY int) {
	Draw_text(positionX+7, positionY+1, "+", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+6, positionY+1, "-", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+5, positionY+1, "-", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+4, positionY+1, "-", termbox.ColorWhite, termbox.ColorDefault)
	Draw_text(positionX+3, positionY+1, "+", termbox.ColorWhite, termbox.ColorDefault)
}

func BottomBeam(positionX, positionY int) {
	Draw_text(positionX+3, positionY+2, "|", termbox.ColorWhite, termbox.ColorDefault)
}

func headHangman(positionX, positionY int) {
	Draw_text(positionX+3, positionY+3, "O", termbox.ColorWhite, termbox.ColorDefault)
}

func bodyHangman(positionX, positionY int) {
	Draw_text(positionX+3, positionY+4, "|", termbox.ColorWhite, termbox.ColorDefault)
}

func leftArmHangman(positionX, positionY int) {
	Draw_text(positionX+2, positionY+4, "/", termbox.ColorWhite, termbox.ColorDefault)
}

func rightArmHangman(positionX, positionY int) {
	Draw_text(positionX+4, positionY+4, "\\", termbox.ColorWhite, termbox.ColorDefault)
}

func leftLegsHangman(positionX, positionY int) {
	Draw_text(positionX+2, positionY+5, "/", termbox.ColorWhite, termbox.ColorDefault)
}

func rightLegsHangman(positionX, positionY int) {
	Draw_text(positionX+4, positionY+5, "\\", termbox.ColorWhite, termbox.ColorDefault)
}
