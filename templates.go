package main

import (
	"html/template"
	"log"
	"net/http"
)

const (
	layout_tmpl string = "./tmpl/layout.html"
	edit_tmpl   string = "./tmpl/edit.html"
	view_tmpl   string = "./tmpl/view.html"
)

var (
	templates = make(map[string]*template.Template)
)

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	log.Println("rendering...")
	log.Println(tmpl)
	err := templates[tmpl].ExecuteTemplate(w, "layout", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
