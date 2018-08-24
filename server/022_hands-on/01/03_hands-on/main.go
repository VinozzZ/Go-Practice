package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

var tpl *template.Template

type User struct {
	Name  string
	Motto string
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "home")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy wow")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	me := User{
		"Yingrong",
		"Eat all the food.",
	}
	// io.WriteString(w, "My name is Yingrong")
	err := tpl.Execute(w, me)
	if err != nil {
		log.Fatalln(err)
	}
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)
}
