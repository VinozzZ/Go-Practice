package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<img src="/toby.jpg">`)
}

/*
http.ServeContent
*/
// func dogPic(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("I'm running")
// 	f, err := os.Open("toby.jpg")
// 	if err != nil {
// 		http.Error(w, "file not found", 404)
// 		return
// 	}
// 	defer f.Close()
// 	fi, err := f.Stat()
// 	if err != nil {
// 		http.Error(w, "file not found", 404)
// 		return
// 	}
// 	http.ServeContent(w, r, f.Name(), fi.ModTime(), f)
// }

/*
http.ServeFile
*/
func dogPic(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "toby.jpg")
}
