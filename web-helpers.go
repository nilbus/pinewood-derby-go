package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"unicode"

	"github.com/shaoshing/train"
)

func ServeWeb() {
	configureTrain()
	configureRoutes()
	err := http.ListenAndServe(":"+port(), nil)
	log.Fatal(err)
}

type restfulHandlerFunc func(http.ResponseWriter, *http.Request, string)

func restfulResource(resourceName string) {
	routeRestfulRequest := func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		resource := path[len("/"+resourceName):len(path)]
		if len(resource) == 0 {
			resource = "/"
		}
		resource = r.Method + " " + resource
		action := "404"
		var id string
		switch {
		case resource == "GET /":
			if path[len(path)-1:len(path)] == "/" {
				http.Redirect(w, r, path[0:len("/"+resourceName)], http.StatusFound)
				return
			}
			action = "Index"
		case resource == "POST /":
			action = "Create"
		case strings.HasPrefix(resource, "GET /"):
			id = resource[len("GET /"):len(resource)]
			switch {
			case id == "new":
				action = "New"
			case strings.HasSuffix(id, "/edit"):
				id = id[0 : len(id)-len("/edit")]
				action = "Edit"
			default:
				action = "Show"
			}
		case strings.HasPrefix(resource, "POST /"):
			id = resource[len("POST /"):len(resource)]
			action = "Update"
		case strings.HasPrefix(resource, "DELETE /"):
			id = resource[len("DELETE /"):len(resource)]
			action = "Delete"
		}
		handler := restfulHandlers["handle"+capitalize(resourceName)+action]
		if handler == nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 Not Found")
		} else {
			handler(w, r, id)
		}
	}
	http.HandleFunc("/"+resourceName, routeRestfulRequest)
	http.HandleFunc("/"+resourceName+"/", routeRestfulRequest)
}

func render(templateName string, w http.ResponseWriter, data interface{}) {
	funcs := template.FuncMap{
		"javascript_tag":  train.JavascriptTag,
		"stylesheet_tag":  train.StylesheetTag,
		"admin":           func() bool { return true },
		"derbyConfig":     func(key string) interface{} { return derbyConfig[key] },
		"eachLane":        eachLane,
		"laneColumnWidth": laneColumnWidth,
		"dashboardJson":   DashboardJson,
	}
	tmpl, err := template.New("layout.tmpl").Funcs(funcs).ParseFiles("templates/layout.tmpl", "templates/"+templateName+".tmpl")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func eachLane() []int {
	lanes := []int{}
	for i := 1; i <= derbyConfig["laneCount"].(int); i++ {
		lanes = append(lanes, i)
	}
	return lanes
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

func capitalize(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func columnWidthSplitInto(columnCount int) int {
	layoutColumnCount := 12
	return layoutColumnCount / columnCount
}

func laneColumnWidth() string {
	return fmt.Sprintf("span%d", columnWidthSplitInto(derbyConfig["laneCount"].(int)))
}
