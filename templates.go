package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var (
	layout_tmpl, _ = filepath.Abs("./tmpl/layout.html")
	edit_tmpl, _   = filepath.Abs("./tmpl/edit.html")
	view_tmpl, _   = filepath.Abs("./tmpl/view.html")
	templates      = make(map[string]*template.Template)
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
