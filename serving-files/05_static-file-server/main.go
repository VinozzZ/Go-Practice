package main

import (
	"net/http"
)

func main() {
	// http.Handle("/", http.FileServer(http.Dir(".")))
	// http.HandleFunc("/dog", dog)
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}

// func dog(w http.ResponseWriter, r *http.Request) {
// w.Header().Set("Content-Type", "text/html; charset=utf-8")
// io.WriteString(w, `<img src="/toby.jpg>"`)
// }
