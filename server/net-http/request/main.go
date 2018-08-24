package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Hotdog int

var tpl *template.Template

func (d Hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}
	data := struct {
		Method        string
		URL           *url.URL
		Submission    url.Values
		Header        http.Header
		Host          string
		ContentLength int64
	}{
		r.Method,
		r.URL,
		r.Form,
		r.Header,
		r.Host,
		r.ContentLength,
	}
	fmt.Println(data)
	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}
func main() {
	var d Hotdog
	http.ListenAndServe(":8080", d)
}
