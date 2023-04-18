package groupie

import (
	"fmt"
	"net/http"
)

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func contains2(w http.ResponseWriter, r *http.Request, slice string, item string) bool {
	r.ParseForm()
	trick := r.Form[slice]
	set := make(map[string]struct{}, len(trick))
	for _, s := range trick {
		set[s] = struct{}{}
	}
	_, ok := set[item]
	return ok
}

func listbox(w http.ResponseWriter, r *http.Request, slice string) string {
	r.ParseForm()
	trick := r.Form[slice]
	set := make(map[string]struct{}, len(trick))
	for _, s := range trick {
		set[s] = struct{}{}
		return s
	}
	return ""
}

func ProcessCheckboxes(w http.ResponseWriter, r *http.Request, Check string, item string) {
	r.ParseForm()
	fmt.Printf("%+v\n", r.Form)
	productsSelected := r.Form[Check]
	fmt.Println(contains(productsSelected, item))
}

func last4(s string) string {
	if len(s) < 4 {
		return ""
	}
	ld := s[len(s)-4:]
	if len(ld) == 4 {
		return ld
	}
	return ""
}

func CutCountry(s string) string {
	var s2 string
	for a := 0; a < len(s); a++ {
		if s[a] != '-' {
			s2 += string(s[a])
		} else {
			break
		}
	}
	return s2
}
