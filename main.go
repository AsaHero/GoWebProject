package main

import (
	"html/template"
	"net/http"
	"os"
)

type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + "txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + "txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}, nil
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/index.html")
	tmp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", nil)
}
