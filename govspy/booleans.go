package main

import (
	"fmt"
	"strconv"
)
func str2int(s string) int{
	i, err := strconv.Atoi(s)
	if err != nil{
		panic("not a number")
	}
	return i
}
func main(){
	fmt.Println(true && false)
	x :=1
	if x != 0{
		fmt.Println("Yes")
	}
	var y []string
	if len(y) !=0{
		fmt.Println("this will not be printed")
	}
}