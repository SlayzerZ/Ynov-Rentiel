package hangman

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Player struct {
	id       string
	email    string
	username string
	password string
}

type HangmanData struct {
	Id			  string
	GamePlayed    string
	VictoryNumber string
	LoseNumber    string
	VictAffNumber string
	Ratio         string
	WorldWin	  string
	WorldLose	  string
}

func ReadCsv(file string) {

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
			id:       line[0],
			email:    line[1],
			username: line[2],
			password: line[3],
		}
		fmt.Println(play.email + " " + play.username + " " + play.password)
	}
}

func Returncsvname(file string, username string) []string {
	name := []string{}
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
			id: 	  line[0],
			email:    line[1],
			username: line[2],
			password: line[3],
		}
		if username == play.username {
			name = []string{play.id, play.email, play.username, play.password}
		}
	}
	return name
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
			id: 	  line[0],
			email:    line[1],
			username: line[2],
			password: line[3],
		}
		if username == play.username && password == play.password {
			fmt.Println(play.username + " " + play.password)
			return true
		}
	}
	return false
}

func ReturnPlayerID(file string, username string, password string) string {
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
			id: 	  line[0],
			email:    line[1],
			username: line[2],
			password: line[3],
		}
		if username == play.username && password == play.password {
			fmt.Println(play.username + " " + play.password)
			return play.id
		}
	}
	return "-1"
}

func ReturnPlayerStats(path string, idUser string) HangmanData {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	hangmandataLast := HangmanData{
		Id: "-1",
		GamePlayed: "0",
		VictoryNumber:"0",
		LoseNumber: "0",
		VictAffNumber: "0",
		Ratio:"0",
	}

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		hangmandata := HangmanData{
			Id: 			line[0],
			GamePlayed: 	line[1],
			VictoryNumber:  line[2],
			LoseNumber: 	line[3],
			VictAffNumber:  line[4],
			Ratio: 			line[5],
		}

		// User Found :
		if hangmandata.Id == idUser {
			hangmandataLast = HangmanData{
				Id: 			hangmandata.Id,
				GamePlayed: 	hangmandata.GamePlayed,
				VictoryNumber:  hangmandata.VictoryNumber,
				LoseNumber: 	hangmandata.LoseNumber,
				VictAffNumber:  hangmandata.VictAffNumber,
				Ratio: 			hangmandata.Ratio,
			}
		}
	}

	// Return Data
	if hangmandataLast.Id != "-1" {
		return hangmandataLast
	} else {
		return hangmandataLast
	}
}

func ReturnWorldStats(path string) (string, string) {
	csvFile, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	var win string
	var lose string

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		win = line[0]
		lose = line[1]
	}

	return win, lose
}

func EditStatsOfPLayer(filePlayer string, fileStats string, idUser string, VictoryNumber int, LoseNumber int, Ratio int, GamePLayed int) string {
	csvFile, err := os.Open(fileStats)
	if err != nil {
		fmt.Println(err)
	}
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}
	for _, line := range csvLines {
		hangmandata := HangmanData{
			Id: 			line[0],
			GamePlayed: 	line[1],
			VictoryNumber:  line[2],
			LoseNumber: 	line[3],
			VictAffNumber:  line[4],
			Ratio: 			line[5],
		}

		if hangmandata.Id == idUser {
			AddStats(filePlayer, fileStats, VictoryNumber, LoseNumber, Ratio, GamePLayed, idUser)
			return "1"
		}
	}

	return "-1"
}

func AddStats(pathPlayer string, pathStats string, VictoryNumber int, LoseNumber int, Ratio int, GamePLayed int, idUser string) {
	fmt.Println("=>=>=>=>=> : ", VictoryNumber, LoseNumber, Ratio, GamePLayed)
	// == Open player.csv ==
	f, err := os.OpenFile(pathPlayer, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// == Open stats.csv ==
	g, err := os.OpenFile(pathStats, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// Create the new stats of the player = new line
	var newStats = []string{idUser, strconv.Itoa(GamePLayed), strconv.Itoa(VictoryNumber), strconv.Itoa(LoseNumber), strconv.Itoa(0), strconv.Itoa(Ratio)}

	var stats [][]string
	stats = append(stats, newStats)

	x := csv.NewWriter(g)
	x.WriteAll(stats)

	if err := x.Error(); err != nil {
		log.Fatal(err)
	}
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
			id: 	  line[0],
			email:    line[1],
			username: line[2],
			password: line[3],
		}
		if mail == play.email || username == play.username {
			fmt.Println("error")
			return false
		} else {
			fmt.Println(play.email + " " + play.username + " " + play.password)
		}
	}
	return true
}

func Addlines(pathPlayer string, pathStats string, playerData []string) {
	// == Open player.csv ==
	f, err := os.OpenFile(pathPlayer, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// == Open stats.csv ==
	g, err := os.OpenFile(pathStats, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// Create the new stats of the player
	var newStats = []string{strconv.Itoa(Count(pathPlayer)), strconv.Itoa(0), strconv.Itoa(0), strconv.Itoa(0), strconv.Itoa(0), strconv.Itoa(0)}

	var stats [][]string
	stats = append(stats, newStats)

	// Get the new Id for the new Player :
	playerData[0] = strconv.Itoa(Count(pathPlayer))

	var data [][]string
	data = append(data, playerData)


	w := csv.NewWriter(f)
	w.WriteAll(data)

	x := csv.NewWriter(g)
	x.WriteAll(stats)

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	if err := x.Error(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Appending success")
}


func AddSingleLines(pathStats string, worldData []string) {
	// == Open stats.csv ==
	g, err := os.OpenFile(pathStats, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	var data [][]string
	data = append(data, worldData)

	x := csv.NewWriter(g)
	x.WriteAll(data)

	if err := x.Error(); err != nil {
		log.Fatal(err)
	}
}
