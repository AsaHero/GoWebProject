package main

import (
	"fmt"
	"html/template"
	"net/http"
)

<<<<<<< HEAD
=======
type User struct {
	Name    string
	Age     int16
	Balance float64
	Hobbies []string
}

>>>>>>> b2710e1db68cb14aa15c2fff8e75ccf3473fedcc
func mainPage(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Printf("Error on parcing tamplate: error %v", err)
	}
	data := User{"Asadbek", 19, 150.12, []string{"Coding", "Music", "BeatMaking"}}
	err = tmp.Execute(w, data)
	if err != nil {
		fmt.Printf("Error on executing template: error %v", err)
	}
}

func main() {
	http.HandleFunc("/", mainPage)
	http.ListenAndServe(":8080", nil)
}
