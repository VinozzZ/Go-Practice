package main

import (
	"fmt"
	"net/http"
)

var hotdog http.Handler

func (hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Mcleod-Key", "this is from mcleod")
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintln(w, "<h1>Writer in the response</h1>")
}

func main() {
	var d Hotdog
	http.ListenAndServe(":8080", d)
}
