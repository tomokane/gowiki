package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"regexp"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := "./data/" + p.Title + ".txt"
	fmt.Println("save()")
	fmt.Println(filename)
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func (p *Page) HTMLData() template.HTML {
	return template.HTML(linking(p.Body))
}

func loadPage(title string) (*Page, error) {
	filename := "./data/" + title + ".txt"
	fmt.Println("loadPage()")
	fmt.Println(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

var linktag = regexp.MustCompile("\\[(.+?)\\]")

func linking(body []byte) []byte {
	return linktag.ReplaceAllFunc(body, func(s []byte) []byte {
		tag := append([]byte("<a href=\""), s[1:len(s)-1]...)
		tag = append(tag, []byte("\">")...)
		tag = append(tag, s[1:len(s)-1]...)
		tag = append(tag, []byte("</a>")...)
		return tag
	})
}
