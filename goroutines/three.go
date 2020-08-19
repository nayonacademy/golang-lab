package main

import (
	"fmt"
	"time"
)

func myAsyncFunction() <-chan int32{
	r := make(chan int32)
	go func() {
		defer close(r)
		time.Sleep(time.Second * 2)
		r <- 2
	}()
	return r
}

func main()  {
	r := <- myAsyncFunction()
	fmt.Println(r)
}

//https://madeddu.xyz/posts/go-async-await/