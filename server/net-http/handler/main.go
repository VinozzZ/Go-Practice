package main

import (
	"fmt"
	"net/http"
)

type Handler int

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "My first server!")
}
func main() {
	var handler Handler
	http.ListenAndServe(":8080", handler)
}
