package main

import (
	"io"
	"net/http"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "home")
}

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Doggy wow")
}

func meHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "My name is Yingrong")
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/me/", meHandler)

	http.ListenAndServe(":8080", nil)
}
