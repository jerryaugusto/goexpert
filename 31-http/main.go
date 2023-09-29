package main

import "net/http"

func main() {
	http.HandleFunc("/", CEPSearch)
	http.ListenAndServe(":8080", nil)
}

func CEPSearch(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, world!"))
}
