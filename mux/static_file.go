package main

import (
	"flag"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "The directory to serve files from. Defaul to the current dir")
	flag.Parse()
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(dir))))
	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// good practice : enforce timeouts for servers you create
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
