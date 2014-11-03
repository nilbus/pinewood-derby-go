package main

import (
	"fmt"
	"net/http"
)

func configureRoutes() {
	http.HandleFunc("/", handleWelcome)
	http.HandleFunc("/board", handleDashboard)
	restfulResource("contestants")
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	render("welcome", w, nil)
}

func handleDashboard(w http.ResponseWriter, r *http.Request) {
	render("dashboard", w, nil)
}

var restfulHandlers = map[string]restfulHandlerFunc{
	"handleContestantsIndex":  handleContestantsIndex,
	"handleContestantsNew":    handleContestantsNew,
	"handleContestantsCreate": handleContestantsCreate,
	"handleContestantsShow":   handleContestantsShow,
	"handleContestantsEdit":   handleContestantsEdit,
	"handleContestantsUpdate": handleContestantsUpdate,
	"handleContestantsDelete": handleContestantsDelete,
}

func handleContestantsIndex(w http.ResponseWriter, r *http.Request, id string) {
	var contestants []Contestant
	db.Find(&contestants)
	render("contestants/index", w, contestants)
}

func handleContestantsNew(w http.ResponseWriter, r *http.Request, id string) {
	render("contestants/new", w, nil)
}

func handleContestantsCreate(w http.ResponseWriter, r *http.Request, id string) {
	contestantName := r.PostFormValue("name")
	contestant := Contestant{Name: contestantName}
	db.Create(&contestant)
	render("contestants/new", w, "Registered "+contestantName)
}

func handleContestantsShow(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "contestants show "+id)
}

func handleContestantsEdit(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "contestants edit "+id)
}

func handleContestantsUpdate(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "contestants update "+id)
}

func handleContestantsDelete(w http.ResponseWriter, r *http.Request, id string) {
	fmt.Fprint(w, "contestants delete "+id)
}
