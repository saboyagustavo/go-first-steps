package main

import "net/http"

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(":8080", nil)
}
