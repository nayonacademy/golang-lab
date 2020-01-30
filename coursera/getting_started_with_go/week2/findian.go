package main

import (
	"fmt"
	"strings"
)
func check(myvalue string) bool{
	if strings.Contains(myvalue, "i") && strings.Contains(myvalue, "a") && strings.Contains(myvalue, "n"){
		return true
	}
	return false
}
func main(){
	var myvalue string
	fmt.Scan(&myvalue)
	result := check(myvalue)
	if result == true{
		fmt.Println("Found!")
	}else {
		fmt.Println("Not Found!")
	}
}
