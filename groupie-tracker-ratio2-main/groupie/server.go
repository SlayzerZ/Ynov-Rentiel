package groupie

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"text/template"
	"time"
)

type Mip struct {
	Names []string
}

type Search struct {
	Name       string
	Scroll     string
	Carrer     string
	Album      string
	Rt         int
	Rt2        int
	Checkbox1  bool
	Checkbox2  bool
	Checkbox3  bool
	Checkbox4  bool
	Checkbox5  bool
	Checkbox6  bool
	Checkbox7  bool
	Checkbox8  bool
	Checkbox9  bool
	Checkbox10 bool
}

type GResults struct {
	Result []Results `json:"results"`
}

type Results struct {
	//Formatted string     `json:"formatted"`
	Geometry Coordonees `json:"geometry"`
}

type Coordonees struct {
	Name      string
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lng"`
}

type Artist struct {
	Id           int      `json:"id"`
	Name         string   `json:"name"`
	Image        string   `json:"image"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Relations    map[string][]string
}

type Relation struct {
	Id             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type WebS struct {
	IdS            string
	Cellphone      string
	Username       string
	Password       string
	Passwordretype string
	IdI            string
	ErrorS         int
}

type WebLog struct {
	IdS       string
	Username  string
	Cellphone string
	Password  string
	IdI       string
	IdI2      int
	ErrorL    int
	ErrorM    int
}

func RemoveArtist(s []Artist, index int) []Artist {
	return append(s[:index], s[index+1:]...)
}

func Main2(w http.ResponseWriter, r *http.Request, pta *Search) {
	pta.Rt2 = 0
	http.Redirect(w, r, "/main", http.StatusSeeOther)
}

func Disconnect(w http.ResponseWriter, r *http.Request, ptn *WebLog) {
	ptn.Username = ""
	ptn.Password = ""
	ptn.Cellphone = ""
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func LogUser(w http.ResponseWriter, r *http.Request, pt *WebLog) {
	pt.Username = r.FormValue("username")
	pt.Password = r.FormValue("password")
	fmt.Println(pt.Username, pt.Password)
	checkl := Checklog("./groupie/csv/users.csv", pt.Username, pt.Password)
	if checkl == false {
		pt.ErrorL = 1
		pt.ErrorM = 1
	} else {
		pt.IdS = ReturnIdS("./groupie/csv/users.csv", pt.Username)
		pt.IdI = ReturnIdI("./groupie/csv/users.csv", pt.Username)
		pt.Cellphone = ReturnCellphone("./groupie/csv/users.csv", pt.Username)
		pt.Username = ReturnUsername("./groupie/csv/users.csv", pt.Username)
		pt.IdI2 = Atoi(pt.IdI)
		pt.ErrorL = 0
		pt.ErrorM = 0
		fmt.Println("IdI : " + pt.IdI)
	}
}

func RegUser(w http.ResponseWriter, r *http.Request, ptr *WebS) {
	rand.Seed(time.Now().UnixNano())
	re := rand.Intn(9) + 1
	// logic part of sign up
	ptr.IdS = strconv.Itoa(Count("./groupie/csv/users.csv"))
	if ptr.IdS == "0" {
		ptr.IdS = "1"
	}
	ptr.Username = r.FormValue("username")
	ptr.Cellphone = r.FormValue("cellphone")
	ptr.Password = r.FormValue("password")
	ptr.Passwordretype = r.FormValue("passwordretype")
	ptr.IdI = strconv.Itoa(re)
	userdata := []string{ptr.IdS, ptr.Username, ptr.Cellphone, ptr.Password, ptr.IdI}
	if ptr.Password != ptr.Passwordretype {
		ptr.ErrorS = 1
	} else {
		checkr := Checkreg("./groupie/csv/users.csv", ptr.Username, ptr.Cellphone)
		if checkr == false {
			ptr.ErrorS = 1
		} else {
			ptr.ErrorS = 0
			Addlines("./groupie/csv/users.csv", userdata)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func signup(w http.ResponseWriter, r *http.Request, ptr *WebS) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		e, _ := template.ParseFiles("templates/register.html", "templates/errorMsg2.html")
		e.Execute(w, nil)
	} else {
		r.ParseForm()
		RegUser(w, r, ptr)
		if ptr.ErrorS == 1 {
			t, _ := template.ParseFiles("templates/register.html", "templates/errorMsg2.html")
			t.Execute(w, ptr)
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}
}

func login(w http.ResponseWriter, r *http.Request, pt *WebLog) {
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("templates/login.html", "templates/errorMsg.html")
		t.Execute(w, pt)
	} else {
		r.ParseForm()
		// logic part of log in
		LogUser(w, r, pt)
		if pt.ErrorL == 1 {
			t, _ := template.ParseFiles("templates/login.html", "templates/errorMsg.html")
			t.Execute(w, pt)
		} else {
			http.Redirect(w, r, "/main", http.StatusSeeOther)
		}
	}
}

func About(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/about.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}

func Profil(w http.ResponseWriter, r *http.Request, pt *WebLog) {
	if pt.ErrorM == 0 {
		pt.IdI = ReturnIdI("./groupie/csv/users.csv", pt.Username)
		pt.IdI2 = Atoi(pt.IdI)
		template, err := template.ParseFiles("templates/profil.html", "templates/pimg.html")
		if err != nil {
			log.Fatal(err)
		}
		template.Execute(w, pt)
	} else {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func Map(w http.ResponseWriter, r *http.Request, pi *Mip) {
	res := GetApi("https://groupietrackers.herokuapp.com/api/artists")
	var artists []Artist
	json.Unmarshal(res, &artists)
	for _, val := range artists {
		pi.Names = append(pi.Names, val.Name)
	}
	template, err := template.ParseFiles("templates/map.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, artists)
}

func Mapartist(w http.ResponseWriter, r *http.Request) {
	co := GetApi("https://api.opencagedata.com/geocode/v1/json?q=" + CutCountry(r.FormValue("q")) + "&key=25eaabf3583443ce9db22e5a5dd4c318")
	var test GResults
	json.Unmarshal(co, &test)
	var mapa Coordonees
	mapa.Name = CutCountry(r.FormValue("q"))
	mapa.Latitude = test.Result[0].Geometry.Latitude
	mapa.Longitude = test.Result[0].Geometry.Longitude
	if mapa.Name == "san_isidro" {
		mapa.Latitude = test.Result[1].Geometry.Latitude
		mapa.Longitude = test.Result[1].Geometry.Longitude
	}
	fmt.Println(mapa)
	template, err := template.ParseFiles("templates/mapartist.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, mapa)
}

func errorPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/errorConnection.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, r)
}

func welcomePage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/welcomepage.html")
	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, r)
}

func mapartist(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("templates/mapartist.html")
	if err != nil {
		log.Fatal(err)
	}

	template.Execute(w, r)
}

func principalPage(w http.ResponseWriter, r *http.Request, pt *WebLog, ptr *Search) {
	res := GetApi("https://groupietrackers.herokuapp.com/api/artists")
	var artists []Artist
	json.Unmarshal(res, &artists)
	var artists2 []Artist
	var artists3 []Artist
	ptr.Checkbox8 = contains2(w, r, "filter", "Carrer")
	ptr.Checkbox9 = contains2(w, r, "filter", "Album")
	ptr.Checkbox10 = contains2(w, r, "filter", "Member")
	if ptr.Checkbox8 {
		ptr.Carrer = r.FormValue("tick")
	} else {
		ptr.Carrer = ""
	}
	if ptr.Checkbox9 {
		ptr.Album = r.FormValue("tick2")
	} else {
		ptr.Album = ""
	}

	if ptr.Rt2 == 0 {
		ptr.Checkbox7 = true
	} else {
		ptr.Checkbox1 = contains2(w, r, "scales", "1")
		ptr.Checkbox2 = contains2(w, r, "scales", "2")
		ptr.Checkbox3 = contains2(w, r, "scales", "3")
		ptr.Checkbox4 = contains2(w, r, "scales", "4")
		ptr.Checkbox5 = contains2(w, r, "scales", "5")
		ptr.Checkbox6 = contains2(w, r, "scales", "6")
		ptr.Checkbox7 = contains2(w, r, "scales", "7+")
		fmt.Println(ptr.Checkbox1, ptr.Checkbox2, ptr.Checkbox3, ptr.Checkbox4, ptr.Checkbox5, ptr.Checkbox6, ptr.Checkbox7)
		fmt.Println(ptr.Checkbox8, ptr.Checkbox9)
	}
	if ptr.Checkbox10 {

		if ptr.Checkbox1 {
			for _, val := range artists {
				if len(val.Members) == 1 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox2 {
			for _, val := range artists {
				if len(val.Members) == 2 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox3 {
			for _, val := range artists {
				if len(val.Members) == 3 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox4 {
			for _, val := range artists {
				if len(val.Members) == 4 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox5 {
			for _, val := range artists {
				if len(val.Members) == 5 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox6 {
			for _, val := range artists {
				if len(val.Members) == 6 {
					artists2 = append(artists2, val)
				}
			}
		}
		if ptr.Checkbox7 {
			for _, val := range artists {
				if len(val.Members) >= 7 {
					artists2 = append(artists2, val)
				}
			}
		}
	} else {
		artists2 = artists
	}
	if ptr.Carrer != "" {
		for _, val := range artists2 {
			if Atoi(ptr.Carrer) >= val.CreationDate {
				artists3 = append(artists3, val)
			}
		}
	} else if ptr.Album != "" {
		for _, val := range artists2 {
			if Atoi(ptr.Album) >= Atoi(last4(val.FirstAlbum)) {
				artists3 = append(artists3, val)
			}
		}
	} else {
		artists3 = artists2
	}

	sort.SliceStable(artists2, func(i, j int) bool {
		return artists2[i].Id < artists2[j].Id
	})
	ptr.Name = r.FormValue("Artiste")
	if pt.ErrorM == 0 {
		if ptr.Name != "" {
			ptr.Rt2 = 0
			for _, val := range artists {
				if ptr.Name == val.Name {
					ptr.Rt = retApi(ptr.Name)
					fmt.Println(ptr.Rt)
					http.Redirect(w, r, "/artist", http.StatusFound)
					return
				}
			}
		} else {
			ptr.Rt2++
		}
		fmt.Println(ptr.Carrer, ptr.Album)
		template, err := template.ParseFiles("templates/principalpage.html")
		if err != nil {
			log.Fatal(err)
		}
		template.Execute(w, artists3)
	} else {
		http.Redirect(w, r, "/error", http.StatusSeeOther)
	}
}

func artistPage(w http.ResponseWriter, r *http.Request, ptr *Search) {
	res := GetApi("https://groupietrackers.herokuapp.com/api/artists/" + r.FormValue("id"))
	reu := GetApi("https://groupietrackers.herokuapp.com/api/relation/" + r.FormValue("id"))
	res2 := GetApi("https://groupietrackers.herokuapp.com/api/artists/" + strconv.Itoa(ptr.Rt))
	reu2 := GetApi("https://groupietrackers.herokuapp.com/api/relation/" + strconv.Itoa(ptr.Rt))
	var artist Artist
	var relation Relation
	if r.FormValue("id") != "" {
		json.Unmarshal(res, &artist)
		json.Unmarshal(reu, &relation)
	} else {
		json.Unmarshal(res2, &artist)
		json.Unmarshal(reu2, &relation)
	}
	artist.Relations = relation.DatesLocations
	template, err := template.ParseFiles("templates/artistpage.html")
	if err != nil {
		log.Fatal(err)
	}
	template.Execute(w, artist)
}

func Execute() {
	// res := GetApi("https://groupietrackers.herokuapp.com/api/artists")
	// var artist []Artist
	// json.Unmarshal(res, &artist)
	// for _, val := range artist {
	// 	fmt.Println((val.CreationDate))
	// }

	datai := Mip{}
	Pti := &datai

	dataU := WebS{"", "", "", "", "", "", 0}
	PtsU := &dataU

	dataL := WebLog{"", "", "", "", "", 0, 0, 1}
	PtL := &dataL

	dataS := Search{"", "artist", "0", "0", 0, 0, false, false, false, false, false, false, true, false, false, false}
	PtS := &dataS

	fmt.Println("http://localhost:8080/")

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		welcomePage(rw, r)
	})
	http.HandleFunc("/login", func(rw http.ResponseWriter, r *http.Request) {
		login(rw, r, PtL)
	})
	http.HandleFunc("/register", func(rw http.ResponseWriter, r *http.Request) {
		signup(rw, r, PtsU)
	})
	http.HandleFunc("/main", func(rw http.ResponseWriter, r *http.Request) {
		principalPage(rw, r, PtL, PtS)
	})
	http.HandleFunc("/main2", func(rw http.ResponseWriter, r *http.Request) {
		Main2(rw, r, PtS)
	})
	http.HandleFunc("/map", func(rw http.ResponseWriter, r *http.Request) {
		Map(rw, r, Pti)
	})
	http.HandleFunc("/mapartist", func(rw http.ResponseWriter, r *http.Request) {
		Mapartist(rw, r)
	})
	http.HandleFunc("/profil", func(rw http.ResponseWriter, r *http.Request) {
		Profil(rw, r, PtL)
	})
	http.HandleFunc("/about", func(rw http.ResponseWriter, r *http.Request) {
		About(rw, r)
	})
	http.HandleFunc("/artist", func(rw http.ResponseWriter, r *http.Request) {
		artistPage(rw, r, PtS)
	})
	http.HandleFunc("/error", func(rw http.ResponseWriter, r *http.Request) {
		errorPage(rw, r)
	})

	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	fi := http.FileServer(http.Dir("./assets/"))
	http.Handle("/assets/", http.StripPrefix("/assets/", fi))
	http.ListenAndServe(":8080", nil)
}
