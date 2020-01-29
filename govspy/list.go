package main

import "fmt"

func main(){
	var numbers []int
	for i:= 1; i<5; i++{
		numbers = append(numbers, i*2)
	}
	fmt.Println(numbers)
	var things []string
	for y := 1; y<7;y++{
		things = append(things, "hello ")
	}
	fmt.Println(things)
}
