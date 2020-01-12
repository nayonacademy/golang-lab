package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"https://google.com",
		"https://fb.com",
		"https://golang.org",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Link might be down")
		c <- "Might be down I think"
		return
	}
	fmt.Println(link, " Link is Ok !")
}
