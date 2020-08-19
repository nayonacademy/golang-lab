package main

import (
	"fmt"
	"time"
)

func background(){
	ticker := time.NewTicker(5*time.Second)
	for t := range ticker.C{
		fmt.Println("tick tok",t)
	}
}

func main()  {
	fmt.Println("Go Tickers Tutorials")
	go background()
	select {}
}