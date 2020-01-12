package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (p Page) saveData() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0666)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there , I love %s", r.URL.Path[1:])
}

func viewHanderl(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	// p, _ := loadPage(title)
	// fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	fmt.Fprintf(w, "Hi there , I love %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHanderl)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
