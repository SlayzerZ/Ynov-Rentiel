// Main File
// Thomas & Ibrahim

package hangman

import (
	"io/ioutil"
)

func AddLetter(allLetterClick []string, letterClick string) []string {
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
				allLetterClick = append(allLetterClick, letterClick)
			}
		}
	}
	return allLetterClick
}

func Letterfound(Fword string, Dword string, letterClick string, attempt int, hangManPosition int, allLetterClick []string) (string, int, int) {
	TFword := []string{}
	TDword := []string{}

	for a := 0; a < len(Fword); a++ {
		TDword = append(TDword, string(rune(Dword[a])))
		TFword = append(TFword, string(rune(Fword[a])))
	}

	var Dword2 string
	var isFundedFinalword bool = false

	for b := 0; b <= len(Fword)-1; b++ {
		if letterClick == string(Fword[b]) {
			TDword[b] = TFword[b]
			isFundedFinalword = true
		}
	}

	// == Check if letterClick is in the allLetterClick Array ==
	if len(allLetterClick) > 0 {
		for i := 0; i < len(allLetterClick); i++ {
			if allLetterClick[i] == letterClick {
				isFundedFinalword = true
			}
		}
	}

	// Check if the letter put in Input (letterclick) is not a letter of the final word
	if !isFundedFinalword && len(letterClick) > 0 {
		attempt++
		hangManPosition++
	}

	if !isFundedFinalword && len(letterClick) >= 2 {
		attempt++
		hangManPosition++
	}

	// Allows you to write to the file
	for c := 0; c < len(TDword); c++ {
		Dword2 += TDword[c]
		ioutil.WriteFile("hangman/solution/Dword.txt", []byte(Dword2), 0)
	}
	return Dword2, attempt, hangManPosition
}

func CheckVictory(WordFinal string, Wordtaped string, WordtoFind string, Attempt int) int {
    // Win
    if (WordFinal == Wordtaped || WordFinal == WordtoFind) && Attempt <= 10 {
        return 1
    }
    if Attempt >= 10 {
        return -1
    }
    return 0
}
