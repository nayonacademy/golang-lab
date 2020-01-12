package main

import "net/http"

import "fmt"

func main() {
	links := []string{
		"https://google.com",
		"https://nayon.net",
	}
	c := make(chan string)
	for _, link := range links {
		go checklink(link, c)
	}
	fmt.Println(<-c)
}

func checklink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, " Link may be down")
		c <- "Might be down now"
		return
	}
	fmt.Println(link, " is up!")
	c <- "Yup its up "
}
