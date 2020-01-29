package main

import "net/http"

import "github.com/gorilla/mux"

import "log"

func YourHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Gorilla\n"))
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", YourHandler)
	log.Fatal(http.ListenAndServe(":8000", r))
}
