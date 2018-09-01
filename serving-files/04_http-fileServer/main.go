package main

import (
	"io"
	"net/http"
)

/*
http.FileServer
*/

// func main() {
// 	http.Handle("/", http.FileServer(http.Dir(".")))
// 	http.HandleFunc("/dog", dog)
// 	http.ListenAndServe(":8080", nil)
// }

// func dog(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `<img src="toby.jpg">`)
// }

/*
http.StripPrefix - it will strip off the "/resource" from `src="/resource/toby.jpg"`
*/

func main() {
	http.HandleFunc("/", dog)
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir("."))))
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resource/toby.jpg"`)
}
