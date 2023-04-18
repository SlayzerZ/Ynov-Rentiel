package hangman

import (
	"github.com/nsf/termbox-go"
)

// Function draws HangMan by hand without .txt file required 
// func DrawHangManNoFile(model, positionX, positionY int) {
// 	// Draw the hangman according to the model entered in the parameters
// 	if model >= 1 {
// 		baseHangMan(positionX, positionY)
// 		if model >= 2 {
// 			beam(positionX, positionY)
// 			if model >= 3 {
// 				HighBeam(positionX, positionY)
// 				if model >= 4 {
// 					BottomBeam(positionX, positionY)
// 					if model >= 5 {
// 						headHangman(positionX, positionY)
// 						if model >= 6 {
// 							bodyHangman(positionX, positionY)
// 							if model >= 7 {
// 								leftArmHangman(positionX, positionY)
// 								if model >= 8 {
// 									rightArmHangman(positionX, positionY)
// 									if model >= 9 {
// 										leftLegsHangman(positionX, positionY)
// 										if model >= 10 {
// 											rightLegsHangman(positionX, positionY)
// 										}
// 									}
// 								}
// 							}
// 						}
// 					}
// 				}
// 			}
// 		}
// 	}
// }

// Function draws the HangMan according to a file .txt The different positions of the HangMan must have a height of 7 and an empty line between each position
func DrawHangMan(model, positionX, positionY, result int) {
	// result == 0 :> Win
	// result == 1 :> Lose
	if model == 0 {
		return
	}

	for i := 1; i <= 8; i++ {
		var line = Read("hangman/asciiart/hangman.txt", i+(model*8)-8)
		if model >= 6 && model <= 8 {
			Draw_text(positionX+1, positionY+i, line, termbox.ColorLightYellow, termbox.ColorDefault)
		} else if model >= 9 {
			if result == 0 {
				Draw_text(positionX+1, positionY+i, line, termbox.ColorLightGreen, termbox.ColorDefault)
			} else {
				Draw_text(positionX+1, positionY+i, line, termbox.ColorLightRed, termbox.ColorDefault)
			}
		} else {
			Draw_text(positionX+1, positionY+i, line, termbox.ColorWhite, termbox.ColorDefault)
		}
	}
}
