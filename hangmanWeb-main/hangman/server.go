package hangman

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type HangmanGame struct {
	Id				string
	Wordtaped       string
	WordtoFind      string
	WordFinal       string
	Allletter       []string
	Attempt         int
	Hangmanposition int
}

type HangmanS struct {
	Id	    	   string
	Email          string
	Username       string
	Password       string
	Passwordretype string
	ErrorS         int
}

type HangmanLog struct {
	Username string
	Password string
	ErrorL   int
}

func LogUser(w http.ResponseWriter, r *http.Request, pt *HangmanLog) {
	pt.Username = r.FormValue("username")
	pt.Password = r.FormValue("password")
	checkl := Checklog("./hangman/Database/players.csv", pt.Username, pt.Password)
	if checkl == false {
		pt.ErrorL = 1
	} else {
		pt.ErrorL = 0
	}
}

func RegUser(w http.ResponseWriter, r *http.Request, ptr *HangmanS) {
	e := false
	// logic part of sign up
	ptr.Email = r.FormValue("email")
	ptr.Username = r.FormValue("username")
	ptr.Password = r.FormValue("password")
	ptr.Passwordretype = r.FormValue("passwordretype")
	playerdata := []string{ptr.Id, ptr.Email, ptr.Username, ptr.Password}

	for l := 0; l < len(ptr.Email); l++ {
		if string(ptr.Email[l]) == "@" {
			e = true
			if e == true {
				ptr.ErrorS = 0
			} else {
				ptr.ErrorS = 1
			}
		}
	}
	if ptr.Password != ptr.Passwordretype {
		ptr.ErrorS = 1
	} else {
		checkr := Checkreg("./hangman/Database/players.csv", ptr.Email, ptr.Username)
		if checkr == false {
			ptr.ErrorS = 1
		} else {
			ptr.ErrorS = 0
			Addlines("./hangman/Database/players.csv", "./hangman/Database/stats.csv", playerdata)
			http.Redirect(w, r, "/", http.StatusSeeOther)
		}
	}
}

func stats(w http.ResponseWriter, r *http.Request, ptl *HangmanLog, ptg *HangmanData) {
	// Check if user as login as an account :
	if ptl.Username == "" || len(ptl.Username) == 0 || len(ptl.Password) == 0 {
		http.Redirect(w, r, "/no-account", http.StatusSeeOther)
	} else {
		template, err := template.ParseFiles("pages/stats.html", "pages/statsPlayer.html")
		if err != nil {
			log.Fatal(err)
		}

		ptg.WorldWin, ptg.WorldLose = ReturnWorldStats("./hangman/Database/world.csv")
		template.Execute(w,ptg)
	}
}

func signup(w http.ResponseWriter, r *http.Request, ptr *HangmanS) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		e, _ := template.ParseFiles("pages/signup.html", "pages/errorMsg2.html")
		e.Execute(w, nil)
	} else {
		r.ParseForm()
		RegUser(w, r, ptr)
		if ptr.ErrorS == 1 {
			t, _ := template.ParseFiles("pages/signup.html", "pages/errorMsg2.html")
			t.Execute(w, ptr)
		} else {
			http.Redirect(w, r, "/difficulty", http.StatusSeeOther)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request, pt *HangmanLog, ptn *HangmanGame, ptD *HangmanData) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("pages/login.html", "pages/errorMsg.html")
		t.Execute(w, pt)
	} else {
		r.ParseForm()
		// logic part of log in
		LogUser(w, r, pt)

		// Assign the Id of the Player :
		ptn.Id = ReturnPlayerID("./hangman/Database/players.csv", pt.Username, pt.Password)
		ptD.Id = ReturnPlayerID("./hangman/Database/players.csv", pt.Username, pt.Password)


		if pt.ErrorL == 1 {
			t, _ := template.ParseFiles("pages/login.html", "pages/errorMsg.html")
			t.Execute(w, pt)
		} else {
			http.Redirect(w, r, "/difficulty", http.StatusSeeOther)
		}
	}
}

func Peasy(w http.ResponseWriter, r *http.Request, ptn *HangmanGame) {
	Easy()
	Fword := Read("./hangman/solution/Fword.txt", 1)
	Dword := Read("./hangman/solution/Dword.txt", 1)
	ptn.WordtoFind = Dword
	ptn.WordFinal = Fword
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func Pmedium(w http.ResponseWriter, r *http.Request, ptn *HangmanGame) {
	Medium()
	Fword := Read("./hangman/solution/Fword.txt", 1)
	Dword := Read("./hangman/solution/Dword.txt", 1)
	ptn.WordtoFind = Dword
	ptn.WordFinal = Fword
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func Phard(w http.ResponseWriter, r *http.Request, ptn *HangmanGame) {
	Hard()
	Fword := Read("./hangman/solution/Fword.txt", 1)
	Dword := Read("./hangman/hangman/solution/Dword.txt", 1)
	ptn.WordtoFind = Dword
	ptn.WordFinal = Fword
	http.Redirect(w, r, "/game", http.StatusSeeOther)
}

func resetGame(ptn *HangmanGame) {
	ptn.Wordtaped = ""
	ptn.WordtoFind = ""
	ptn.WordFinal = ""
	ptn.Allletter = []string{}
	ptn.Attempt = 0
	ptn.Hangmanposition = 0
}

func Difficulty(w http.ResponseWriter, r *http.Request, ptn *HangmanLog, ptg *HangmanGame) {
	t, err := template.ParseFiles("pages/difficulty.html")
	resetGame(ptg)

	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, ptn)
}

// Game Function :
func Game(w http.ResponseWriter, r *http.Request, ptn *HangmanGame, ptl *HangmanLog, ptD *HangmanData) {
	t, err := template.ParseFiles("pages/game.html", "pages/pendu.html")
	if err != nil {
		log.Fatal(err)
	}
	// Check if user as login as an account :
	if ptl.Username == "" || len(ptl.Username) == 0 || len(ptl.Password) == 0 {
		http.Redirect(w, r, "/no-account", http.StatusSeeOther)
	}

	// Initializ FinalWord and CurrentWord:
	Fword := Read("./hangman/solution/Fword.txt", 1)
	Dword := Read("./hangman/solution/Dword.txt", 1)

	// Get the Letter Input:
	r.ParseForm()
	ptn.Wordtaped = r.FormValue("letterclick")

	// Initialize the *HangManData:
	ptn.WordtoFind = Dword
	ptn.WordFinal = Fword

	// Check if letterclick are in the finalWord:
	// Yes: Change Dword.txt
	// No: attempt++ & hangmanposition++
	var value1, value2, value3 = Letterfound(ptn.WordFinal, ptn.WordtoFind, ptn.Wordtaped, ptn.Attempt, ptn.Hangmanposition, ptn.Allletter)
	ptn.Attempt = value2
	ptn.Hangmanposition = value3
	ptn.WordtoFind = value1

	// If letterclick are not a the finalWord
	// => Add LetterClick into AllLetter Array
	ptn.Allletter = AddLetter(ptn.Allletter, ptn.Wordtaped)

	fmt.Println("=>=>=>=> : HangManGame : ", ptn)


	// Check if the game are finish or not:
	// Player Win :
	if CheckVictory(ptn.WordFinal, ptn.Wordtaped, ptn.WordtoFind, ptn.Attempt) == 1 {
		 
		
		if Atoi(ptD.GamePlayed) == 0 {
			ptD.GamePlayed = "1"
		}

		ptD.VictoryNumber = strconv.Itoa(Atoi(ptD.VictoryNumber) +1)
		ptD.GamePlayed = strconv.Itoa(Atoi(ptD.GamePlayed) +1)

		ptD.Ratio = strconv.Itoa((Atoi(ptD.VictoryNumber) / Atoi(ptD.GamePlayed)) * 100)

		EditStatsOfPLayer("./hangman/Database/players.csv", "./hangman/Database/stats.csv", ptn.Id, Atoi(ptD.VictoryNumber), Atoi(ptD.LoseNumber), Atoi(ptD.Ratio), Atoi(ptD.GamePlayed))
		
		// == Add World Data ==
		var valueWin, valueLose = (ReturnWorldStats("./hangman/Database/world.csv"))

		AddSingleLines("./hangman/Database/world.csv", []string{strconv.Itoa(Atoi(valueWin)+1), strconv.Itoa(Atoi(valueLose))})

		// == Redirect ==
		http.Redirect(w, r, "/victory", http.StatusSeeOther)
	}
	// Player Lose :
	if CheckVictory(ptn.WordFinal, ptn.Wordtaped, ptn.WordtoFind, ptn.Attempt) == -1 {
		ptD.VictoryNumber = strconv.Itoa(Atoi(ptD.LoseNumber) +1)
		ptD.GamePlayed = strconv.Itoa(Atoi(ptD.GamePlayed) +1)

		if Atoi(ptD.GamePlayed) == 0 {
			ptD.GamePlayed = "1"
		}

		ptD.Ratio = strconv.Itoa((Atoi(ptD.VictoryNumber) / Atoi(ptD.GamePlayed)) *100)

		EditStatsOfPLayer("./hangman/Database/players.csv", "./hangman/Database/stats.csv", ptn.Id, Atoi(ptD.VictoryNumber), Atoi(ptD.LoseNumber), Atoi(ptD.Ratio), Atoi(ptD.GamePlayed))
		
		// == Add World Data ==
		var valueWin, valueLose = ReturnWorldStats("./hangman/Database/world.csv")
		AddSingleLines("./hangman/Database/world.csv", []string{strconv.Itoa(Atoi(valueWin)), strconv.Itoa(Atoi(valueLose)+1)})
		
		// == Redirect ==
		http.Redirect(w, r, "/lose", http.StatusSeeOther)
	}

	t.Execute(w, ptn)
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("pages/errorConnection.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}

func quatrecentquatre(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("pages/404.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}

// Game finish Win/Lose :
func gameFinish(w http.ResponseWriter, r *http.Request, winOrLose bool, ptn *HangmanGame) {
	// Win :
	if winOrLose {
		template, err := template.ParseFiles("pages/victory.html")
		if err != nil {
			log.Fatal(err)
		}
		template.Execute(w, r)
	} else if !winOrLose && ptn.Attempt >= 10 {
		template, err := template.ParseFiles("pages/lose.html")
		if err != nil {
			log.Fatal(err)
		}
		template.Execute(w, r)
	} else {
		http.Redirect(w, r, "/404", http.StatusSeeOther)
	}
}

func Execute() {
	dataU := HangmanS{"", "", "", "", "", 0}
	PtsU := &dataU

	dataL := HangmanLog{"", "", 0}
	PtL := &dataL

	dataG := HangmanGame{"0", "", Read("./hangman/solution/Dword.txt", 1), "", []string{}, 0, 0}
	PtG := &dataG

	dataD := HangmanData{"0", "0", "0", "0", "0", "0", "0", "0"}
	PtD := &dataD

	fmt.Println("http://localhost:8080/")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		login(rw, r, PtL, PtG, PtD)
	})

	// http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
	// 	login(rw, r, PtL, PtG, PtD)
	// })

	http.HandleFunc("/signup", func(rw http.ResponseWriter, r *http.Request) {
		signup(rw, r, PtsU)
	})

	http.HandleFunc("/stats", func(rw http.ResponseWriter, r *http.Request) {
		dataH := ReturnPlayerStats("./hangman/Database/stats.csv", PtG.Id)
		PtH := &dataH
		
		stats(rw, r, PtL, PtH)
	})

	http.HandleFunc("/difficulty", func(rw http.ResponseWriter, r *http.Request) {
		Difficulty(rw, r, PtL, PtG)
	})

	http.HandleFunc("/game", func(rw http.ResponseWriter, r *http.Request) {
		Game(rw, r, PtG, PtL, PtD)
	})

	http.HandleFunc("/easy", func(rw http.ResponseWriter, r *http.Request) {
		Peasy(rw, r, PtG)
	})

	http.HandleFunc("/medium", func(rw http.ResponseWriter, r *http.Request) {
		Pmedium(rw, r, PtG)
	})

	http.HandleFunc("/hard", func(rw http.ResponseWriter, r *http.Request) {
		Phard(rw, r, PtG)
	})

	http.HandleFunc("/victory", func(rw http.ResponseWriter, r *http.Request) {
		gameFinish(rw, r, true, PtG)
	})

	http.HandleFunc("/lose", func(rw http.ResponseWriter, r *http.Request) {
		gameFinish(rw, r, false, PtG)
	})

	http.HandleFunc("/no-account", func(rw http.ResponseWriter, r *http.Request) {
		errorPage(rw, r)
	})

	http.HandleFunc("/404", func(rw http.ResponseWriter, r *http.Request) {
		quatrecentquatre(rw, r)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./img/"))
	http.Handle("/img/", http.StripPrefix("/img/", fi))
	ff := http.FileServer(http.Dir("./font/"))
	http.Handle("/font/", http.StripPrefix("/font/", ff))
	http.ListenAndServe(":8080", nil)
}
