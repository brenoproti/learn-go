package main

import "net/http"

type Example struct {
	title string
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HomeHandler)
	mux.Handle("/example", Example{title: "example"})
	http.ListenAndServe(":8080", mux)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from mux"))
}

func (e Example) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(e.title))
}
