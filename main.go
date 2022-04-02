package main

import (
	"GoWebProject/data"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/index.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}

	err = tm.ExecuteTemplate(w, "index", data.SelectAll())
	if err != nil {
		panic(err)
	}
}
func new_post(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		r.ParseForm()
		err := data.Insert(r.FormValue("title"), r.FormValue("full_text"))
		if err != nil {
			panic(err)
		}
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	} else {
		tm, err := template.ParseFiles("template/new_post.html", "template/header.html", "template/footer.html")
		if err != nil {
			panic(err)
		}
		err = tm.ExecuteTemplate(w, "new_post", nil)
		if err != nil {
			panic(err)
		}
	}
}
func about(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/about.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w, "about", nil)
	if err != nil {
		panic(err)
	}
}
func contact(w http.ResponseWriter, r *http.Request) {
	tm, err := template.ParseFiles("template/contact.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	err = tm.ExecuteTemplate(w, "contact", nil)
	if err != nil {
		panic(err)
	}
}

func post(w http.ResponseWriter, r *http.Request) {

	i, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	article, err := data.FindByID(uint16(i))
	if err != nil {
		panic(err)
	}

	tm, err := template.ParseFiles("template/post.html", "template/header.html", "template/footer.html")
	if err != nil {
		panic(err)
	}
	tm.ExecuteTemplate(w, "post", &article)
}

func deletePost(w http.ResponseWriter, r *http.Request) {
	i, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		panic(err)
	}
	err = data.DeleteById(uint16(i))
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

func handleFunc() {
	// Static files
	fs := http.FileServer(http.Dir("./static/"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Handling requests
	http.HandleFunc("/delete/post", deletePost)
	http.HandleFunc("/post", post)
	http.HandleFunc("/contact", contact)
	http.HandleFunc("/about", about)
	http.HandleFunc("/new_post", new_post)
	http.HandleFunc("/", index)

	log.Fatalln(http.ListenAndServe(":8080", nil))
}

func main() {
	handleFunc()
}
