package main

import (
	"fmt"
)

func says(s string, c chan int){
	total := 6
	for i:=0; i<5;i++{
		total += i
		fmt.Println(s)
	}
	c <- total
}
func main()  {
	c := make(chan int)
	go says("hello",c)
	x := <-c
	fmt.Println(x)

}