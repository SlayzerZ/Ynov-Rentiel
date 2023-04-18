package hangman

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Settings function: 
func Settings() {
	rand.Seed(time.Now().UnixNano())
	// Takes as command parameter the file containing the words
	sentence := os.Args[1]
	// Randomly selects a word from the file
	r := rand.Intn(Count(sentence)) + 1
	strin := Read(sentence, r)
	// Randomly selects letters from the word
	random := rand.Intn(len(strin) + 1)
	rn := len(strin)/2 - 1
	strinunder := strings.Repeat("_", len(strin))
	Tabstrin := []string{}
	Tabstrinunder := []string{}
	var strinunder2 string
	// Write in the file the final word
	ioutil.WriteFile("hangman/solution/Fword.txt", []byte(strin), 0)

	for a := 0; a < len(strin); a++ {
		Tabstrin = append(Tabstrin, string(rune(strin[a])))
		Tabstrinunder = append(Tabstrinunder, string(rune(strinunder[a])))
	}

	for c := 0; c < rn; c++ {
		if rn > 0 {
			random = rand.Intn(len(Tabstrin))
			Tabstrinunder[random] = Tabstrin[random]
		}
	}

	for b := 0; b < len(Tabstrinunder); b++ {
		strinunder2 += string(Tabstrinunder[b])
	}

	// Write in the file the word with the letters a found
	ioutil.WriteFile("hangman/solution/Dword.txt", []byte(strinunder2), 0)
}
