package main

import "fmt"

func main(){
	elements := make(map[string]int)
	elements["h"] = 1
	elements["o"] = 5
	delete(elements, "o")
	fmt.Println(elements["h"], elements["o"])

	if number, ok := elements["o"]; ok{
		fmt.Println(number)
	}
}