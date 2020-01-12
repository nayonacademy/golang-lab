package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	newColro := make(map[int]string)

	newColro[10] = "#ffffff"
	delete(newColro, 10)
	fmt.Println(colors)
	fmt.Println(newColro)
}
