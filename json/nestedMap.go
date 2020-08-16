package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Printf("Demo 1")
	Demo1()
	Demo2()
}

func Demo1(){
	categories := map[string][]string{
		"cat1": []string{"pro 1", "pro 2"},
		"cat 2": []string{"pro 3", "pro 4"},
	}
	fmt.Println(categories)
	for key, value := range categories{
		fmt.Println(key)
		for _, v := range value{
			fmt.Println("\t", v)
		}
	}
}

func Demo2(){
	categories := make(map[string][]string)
	categories["cat1"] = append(categories["cat1"],"product 1")
	categories["cat1"] = append(categories["cat1"],"product 1")
	categories["cat1"] = append(categories["cat1"],"product 1")
	categories["cat1"] = append(categories["cat1"],"product 1")
	categories["cat2"] = append(categories["cat2"],"product 1")
	fmt.Println(categories)
	fmt.Println(strings.Repeat("#",10))
	for key, value := range categories{
		fmt.Println(key)
		for _, v := range value{
			fmt.Println("\t", v)
		}
	}
}