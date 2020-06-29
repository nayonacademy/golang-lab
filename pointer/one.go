package main

import "fmt"

func main(){
	var x int  = 5748
	var p *int

	p = &x

	var s *int

	fmt.Println("values stored in x = ",x)
	fmt.Println("Address of x = ", &x)
	fmt.Println("value stored in variable p = ", p)
	fmt.Println("s = ",s)
}
