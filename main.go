package main

import (
	"html/template"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("views/index.html")
	tmp.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", nil)
}
