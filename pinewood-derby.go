package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	// "github.com/codegangsta/martini-contrib/binding"
	// "database/sql"
	// _ "github.com/mattn/go-sqlite3"
	// "html/template"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	m.Get("/", func(r render.Render) {
		r.HTML(200, "welcome", nil)
	})

	m.Run()
}
