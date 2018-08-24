package main

import (
	"io"
	"net/http"
)

/*
	Basic way to do routing
*/
// type hotdog int

// func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	switch req.URL.Path {
// 	case "/dog":
// 		io.WriteString(w, "doggy")
// 	case "/cat":
// 		io.WriteString(w, "kitty")
// 	default:
// 		io.WriteString(w, "home")
// 	}
// }

// func main() {
// 	var h hotdog
// 	http.ListenAndServe(":8080", h)
// }

/*
	Use ServeMux for routing
*/

// type hotdog int

// func (m hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "doggy")
// }

// type hotcat int

// func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "kitty")
// }

// func main() {
// 	var h hotdog
// 	var c hotcat

// 	mux := http.NewServeMux()
// 	mux.Handle("/dog/", h)
// 	mux.Handle("/cat", c)
// 	http.ListenAndServe(":8080", mux)

// }

/*
	Use Default ServeMux
*/

func dogHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "doggy")
}

func catHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "kitty")
}

func main() {
	http.HandleFunc("/dog/", dogHandler)
	http.HandleFunc("/cat/", catHandler)
	// When pass in nil to the handler argument, it will use the default serveMux
	http.ListenAndServe(":8080", nil)
}
