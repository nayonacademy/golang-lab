package main

import (
	"fmt"
	"time"
)

func backgroud(s string){
	for i := 0; i < 5 ; i++{
		fmt.Println(s)
		time.Sleep(100 * time.Millisecond)
	}
}

func main(){
	go backgroud("hello")
	backgroud("world")
}