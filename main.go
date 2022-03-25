package main

import (
	// "database/sql"
	// "fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func index(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w,"index", nil)
	if err != nil {
		panic(err)
	}
}
func post(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/post.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w,"post", nil)
	if err != nil {
		panic(err)
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/about.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w,"about", nil)
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/contact.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w,"contact", nil)
	if err != nil {
		panic(err)
	}
}
func handleFunc() {
	// Static files
	fs := http.FileServer(http.Dir("./static/")) 
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handling requests
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/post", post)
	http.HandleFunc("/", index)


	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func main() {
	handleFunc()
}
