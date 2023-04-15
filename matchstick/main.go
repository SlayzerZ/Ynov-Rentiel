package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func IsNumeric(s string) bool {
	for _, letter := range s {
		if int(letter) >= 0o0 && int(letter) <= 47 || int(letter) >= 58 && int(letter) <= 127 {
			return false
		}
	}
	return true
}

func Atoi(s string) int {
	var res int
	trune := []rune(s)
	negatif := 1
	for i, c := range trune {
		if c == '-' {
			if i != 0 {
				return 0
			} else {
				negatif = -1
			}

		} else if c == '+' {
			if i != 0 {
				return 0
			}
		} else if !IsNumeric(string(c)) {
			return 0
		} else {
			if i == len(trune)-1 {
				res += int(c - 48)
			} else {
				res += int(c - 48)
				res *= 10
			}

		}

	}
	return res * negatif
}

func Allumette(nb int) {
	if nb > 15 {
		return
	}
	for i := 1; i <= nb; i++ {
		if i < nb {
			fmt.Print("|")
		} else {
			fmt.Println("|")
		}
	}
}

func Read() string {
	scanner := bufio.NewScanner(os.Stdin)
	for a := 15; a >= 0; a-- {
		scanner.Scan()
		text := scanner.Text()
		return text
	}
	return "n"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3) + 1
	fmt.Println("Choose between 1 or 15 matchstick")
	A := Atoi(Read())
	Allumette(A)
	for t := 1; t < 15; t++ {
		P1t := "Choose: "
		fmt.Print(P1t)
		NA := Atoi(Read())
		if NA <= 0 {
			fmt.Println("Entry must be a number cannot be negative or equal to 0")
			A = A
		} else if NA > 3 {
			fmt.Println("Number must be between 1 and 3")
			A = A
		} else {
			Allumette(A - NA)
			A -= NA
			if A <= 0 {
				fmt.Println("You lost... noob !")
				break
			}
			NA = r
			IAp := "IA remove "
			IAn := strconv.Itoa(r)
			fmt.Println(IAp + IAn + " matchstick")
			Allumette(A - NA)
			A -= NA
			if A <= 0 {
				fmt.Println("You won !")
				break
			}
		}

	}
}
