package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name    string
	Age     int
	Balance float64
	Hobbies []string
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Printf("Error on parcing tamplate: error %v", err)
	}
	data := User{"Nazarbek", 13, 150.12, []string{"Coding", "Music", "BeatMaking"}}
	err = tmp.Execute(w, data)
	if err != nil {
		fmt.Printf("Error on executing template: error %v", err)
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", nil)
}
