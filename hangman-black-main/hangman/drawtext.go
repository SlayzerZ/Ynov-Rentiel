package hangman

import (
	"github.com/nsf/termbox-go"
)

// Function to draw the text
func Draw_text(positionX, positionY int, text string, fontColor, backgroundColor termbox.Attribute) {
	for i := 0; i <= len(text)-1; i++ {
		termbox.SetCell(positionX+i, positionY, rune(text[i]), fontColor, backgroundColor)
	}
}
