package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandlerFunc("/", foo)
	http.Handle("/favicon.io",http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	log.Println(ctx)
	fmt.Fprintln(w, ctx)
}
