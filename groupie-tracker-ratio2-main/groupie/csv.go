package groupie

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Player struct {
	idS       string
	username  string
	cellphone string
	password  string
	idI       string
}

func ReturnIdS(file string, username string) string {

	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.username || username == play.cellphone {
			return play.idS
		}
	}
	return ""
}

func ReturnIdI(file string, username string) string {

	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.username || username == play.cellphone {
			return play.idI
		}
	}
	return ""
}

func ReturnUsername(file string, username string) string {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.cellphone {
			username = play.username
		}
	}
	return username
}

func ReturnCellphone(file string, username string) string {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.username {
			username = play.cellphone
		}
	}
	return username
}

func Checklog(file string, username string, password string) bool {

	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.username || username == play.cellphone {
			if password == play.password {
				fmt.Println(play.username + " " + play.password)
				return true
			}
		}
	}
	return false
}

func Checkreg(file string, mail string, username string) bool {
	csvFile, err := os.Open(file)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		play := Player{
			idS:       line[0],
			username:  line[1],
			cellphone: line[2],
			password:  line[3],
			idI:       line[4],
		}
		if username == play.username {
			fmt.Println("error")
			return false
		} else {
			fmt.Println(play.username + " " + play.cellphone + " " + play.password)
		}
	}
	return true
}

func Addlines(pathPlayer string, playerData []string) {
	// == Open player.csv ==
	f, err := os.OpenFile(pathPlayer, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var data [][]string
	data = append(data, playerData)

	w := csv.NewWriter(f)
	w.WriteAll(data)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending success")
}

func Atoi(s string) int {
	var numberString string
	var numberReverse string
	var number int = 0
	var coef int = 1
	var negativeNumber bool = false
	var first bool = true

	allNumber := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	allNumberString := [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	if len(s) == 0 {
		return 0
	}

	// ==> Loop to convert string number into int number
	for i := 0; i <= len(s)-1; i++ {
		if (s[i] >= 48 && s[i] <= 57) || s[i] == 45 {
			if s[i] == 45 && first == true {
				if i <= len(s)-2 && (s[i+1] == 45 || s[i+1] == 43) {
					return 0
				}
				negativeNumber = true
				first = false
			} else if first == false && (s[i] == 45 || s[i] == 43) {
				return 0
			} else {
				numberString += string(s[i])
				first = false
			}
		} else {
			if s[len(s)-1] == 43 || s[len(s)-1] == 45 {
				return 0
			} else if i <= len(s)-2 && s[0] == 43 && s[1] == 43 {
				return 0
			} else if s[i] != 43 {
				return 0
			}
		}
	}

	// ==> Reversr
	for j := len(numberString) - 1; j >= 0; j-- {
		numberReverse += string(numberString[j])
	}

	for k := 0; k <= len(numberReverse)-1; k++ {
		for l := 0; l <= len(allNumberString)-1; l++ {
			if allNumberString[l] == string(numberReverse[k]) {
				number += (allNumber[l] * coef)
				coef *= 10
			}
		}
	}

	if negativeNumber {
		number = number * (-1)
	}
	return number
}
