package main

import "fmt"

func sum(a int, c chan int){
	c <- a * a
}

func division(a, b int, d chan int){
	d <- b/a
}
func main(){
	c := make(chan int)
	d := make(chan int)
	go division(10,400, d)
	go sum(45, c)
	me := <-c
	you := <-d
	fmt.Println(me)
	fmt.Println(you)
	select {
	case c:= <-c:
		fmt.Println("received",c)
	}
}