// Main File
// Thomas & Ibrahim

package main

import (
	"fmt"
	"hangman/hangman"
	"io/ioutil"

	"github.com/nsf/termbox-go"
)

func main() {
	// Variable Declaration :
	var currentMenu int = 0
	var letterClick string = ""
	var allLetterClick string = ""
	var attempt int
	var hangManPosition int = 0
	var reloadGame bool = true
	var winRate int = 0
	var loseRate int = 0
	var attemptTotal int = 0

	// Init Termbox :
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	// Load / Reload Settings :
	hangman.Settings()

	// Init Final Word & Current Word :
	Fword := hangman.Read("hangman/solution/Fword.txt", 1)
	Dword := hangman.Read("hangman/solution/Dword.txt", 1)

	Dword2 := ""
	TFword := []string{}
	TDword := []string{}

// Name MainLoop :
mainloop:
	for {
		// Reload Game => Reset all Variable :
		if reloadGame {
			letterClick = ""
			allLetterClick = ""
			attempt = 0
			hangManPosition = 0
			hangman.Settings()
			Dword2 = ""
			TFword = []string{}
			TDword = []string{}
			Fword = hangman.Read("hangman/solution/Fword.txt", 1)
			Dword = hangman.Read("hangman/solution/Dword.txt", 1)

			for a := 0; a < len(Fword); a++ {
				TDword = append(TDword, string(rune(Dword[a])))
				TFword = append(TFword, string(rune(Fword[a])))
			}

			reloadGame = false
		}

		// First Page - Welcome
		if currentMenu == 0 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			hangman.Draw_text(0, 0, "\"ESC\" to quit, \"LEFT\" or \"RIGHT\" to switch tabs", termbox.ColorLightMagenta, termbox.ColorDefault)
			hangman.BarMenu(1, 0)
			hangman.Welcome()
			termbox.Flush()
		}

		// Second Page - Game
		if currentMenu == 1 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			hangman.Draw_text(0, 0, "\"ESC\" to quit, \"LEFT\" or \"RIGHT\" to switch tabs", termbox.ColorLightMagenta, termbox.ColorDefault)
			hangman.BarMenu(1, 1)
			hangman.GameDisplay(letterClick, allLetterClick, attempt)
			termbox.Flush()
		}

		// Third Page - Help
		if currentMenu == 2 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			hangman.Draw_text(0, 0, "\"ESC\" to quit, \"LEFT\" or \"RIGHT\" to switch tabs", termbox.ColorLightMagenta, termbox.ColorDefault)
			hangman.BarMenu(1, 2)
			hangman.Help()
			termbox.Flush()
		}

		// Page 4 - Lose
		if currentMenu == 3 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			hangman.Draw_text(0, 0, "\"ESC\" to quit, \"LEFT\" or \"RIGHT\" to switch tabs", termbox.ColorLightMagenta, termbox.ColorDefault)
			hangman.BarMenuResult(1, 1)
			hangman.Lose()
			termbox.Flush()
		}

		// Page 5 - Win
		if currentMenu == 4 {
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			hangman.Draw_text(0, 0, "\"ESC\" to quit, \"LEFT\" or \"RIGHT\" to switch tabs", termbox.ColorLightMagenta, termbox.ColorDefault)
			hangman.BarMenuResult(1, 0)
			hangman.Win()
			termbox.Flush()
		}

		// No Page - Fixe HangMan Draw (Game)
		if currentMenu == 1 {
			for i := 0; i <= 10; i++ {
				if i == hangManPosition {
					hangman.DrawHangMan(i, 42, 6, 1)
				}
			}
			if hangManPosition >= 10 {
				hangman.DrawHangMan(10, 42, 6, 1)
			}
			termbox.Flush()
		}

		// Init KeyBoard PollEvent :
		// Wait for an event and return it.
		ev := termbox.PollEvent()
		// See if the event type is a Keyboard event :
		if ev.Type == termbox.EventKey {
			if ev.Key == termbox.KeyEsc {
				break mainloop
			} else if ev.Key == termbox.KeyArrowRight && attempt <= 9 && currentMenu != 4 {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				currentMenu += 1
				if currentMenu == 3 {
					currentMenu = 0
				}
			} else if ev.Key == termbox.KeyArrowLeft && attempt <= 9 && currentMenu != 4 {
				termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
				currentMenu -= 1
				if currentMenu == -1 {
					currentMenu = 2
				}
			} else {
				// Menu condition is = 1
				if currentMenu == 1 {
					if ev.Key == termbox.KeyEnter {
						if letterClick != "_" {

							var add bool = true
						
							for i := 0; i <= len(allLetterClick)-1; i++ {
								if string(allLetterClick[i]) == letterClick {
									add = false
								}
							}

							if add {
								if len(letterClick) == 0 {
									add = false
								} else {
									if len(letterClick) == 1 {
										allLetterClick += letterClick + ","
									}
								}
							}

							// Word found
							if letterClick == Fword {
								winRate++
								Dword2 = ""
								ioutil.WriteFile("hangman/solution/Dword.txt", []byte(Dword2), 0)
								currentMenu = 4
								continue
							} else {
								if len(letterClick) >= 2 {
									attempt++
									attemptTotal++
									hangManPosition++
								}
							}
							
							// See if the letter in Input is a letter of the final word
							var isFundedFinalword bool = false
							for b := 0; b <= len(Fword)-1; b++ {
								if letterClick == string(Fword[b]) {
									TDword[b] = TFword[b]
									isFundedFinalword = true
								}
							}

							// Check if the letter put in Input (letterclick) is not a letter of the final word
							if !isFundedFinalword && add {
								attempt++
								attemptTotal++
								hangManPosition++
							}

							// Allows you to write to the file
							for c := 0; c < len(TDword); c++ {
								Dword2 += TDword[c]
								ioutil.WriteFile("hangman/solution/Dword.txt", []byte(Dword2), 0)
							}

							// See if the word found is equal to the final word
							if Dword2 == Fword {
								winRate++
								Dword2 = ""
								ioutil.WriteFile("hangman/solution/Dword.txt", []byte(Dword2), 0)
								currentMenu = 4
								continue
							}

							// See if the player has lost the game
							if attempt >= 10 {
								loseRate++
								attempt = 10
								currentMenu = 3
								Dword2= ""
								ioutil.WriteFile("hangman/solution/Dword.txt", []byte(Dword2), 0)
								continue
							}

							// Resets the word searched
							Dword2 = ""
						}
						// Resets the letter to Input
						letterClick = ""

					} else {
						// Add letter into letterClick
						if rune(ev.Ch) != 0 && len(letterClick) <= 18 {
							letterClick += string(ev.Ch)
						} else {
							// Delete last letter of letterClick with index
							if len(letterClick) > 0 {
								letterClick = letterClick[0:len(letterClick)-1]
							}
						}
					}
				}

				// If current menu = lose
				if currentMenu == 3 {
					if ev.Key == termbox.KeyEnter {
						currentMenu = 1
						reloadGame = true
					}
				}

				// If current menu = Win
				if currentMenu == 4 {
					if ev.Key == termbox.KeyEnter {
						currentMenu = 1
						reloadGame = true
					}
				}
			}
		}
	}
	// Stop Game: 
	termbox.Close()

	// Close - Initialize Stats of all Game
	var percentage int

	if winRate == 0 && loseRate == 0 {
		percentage = 0
	} else {
		percentage = int((winRate*100)/(winRate+loseRate))
	}
	// Close - Game Finish - Show Stats of Player
	fmt.Print("\n")
	fmt.Println("==============================")
	fmt.Println(" | => Finished")
	fmt.Println(" | You won :", winRate, "times")
	fmt.Println(" | You lose :", loseRate, "times")
	fmt.Printf(" | Win Percentage : %d%s\n", percentage, "%")
	fmt.Printf(" | Total number of attempts : %d\n", attemptTotal)
	fmt.Println("==============================")
	fmt.Print("\n")
}
