package main

import (
	"fmt"
	"net/http"
)

func mainPage(w http.ResponseWriter, r http.Request) {

}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<a href=\"https://google.com/\"> Hello World </a>")
	})
	http.ListenAndServe(":8080", nil)
}
