package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/shaoshing/train"
)

func ServeWeb() {
	configureTrain()
	configureRoutes()
	err := http.ListenAndServe(":"+port(), nil)
	log.Fatal(err)
}

func configureRoutes() {
	http.HandleFunc("/", handleWelcome)
}

func handleWelcome(w http.ResponseWriter, r *http.Request) {
	funcs := template.FuncMap{
		"javascript_tag": train.JavascriptTag,
		"stylesheet_tag": train.StylesheetTag,
	}
	tmpl, err := template.New("layout.tmpl").Funcs(funcs).ParseFiles("templates/layout.tmpl", "templates/welcome.tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// Return the $PORT environment variable, or "3000" otherwise
func port() string {
	if port := os.Getenv("PORT"); port != "" {
		return port
	} else {
		return "3000"
	}
}

func configureTrain() {
	train.Config.SASS.DebugInfo = true
	train.Config.SASS.LineNumbers = true
	train.Config.Verbose = true
	train.ConfigureHttpHandler(nil)
}
