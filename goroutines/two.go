package main

import "fmt"

func sum(s []int, c chan int){
	sum := 0
	for _, v := range s{
		sum += v
	}
	c <- sum
}

func main()  {
	c := make(chan int)
	a := []int{7, 2, 8, -9, 4, 0}
	go sum(a,c)
	x := <-c
	fmt.Println(x)
}
