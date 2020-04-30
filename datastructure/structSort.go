package main

import (
	"fmt"
	"sort"
)

type Mobile struct {
	Brand string
	Price int
}
type ByPrice []Mobile

func (a ByPrice) Len() int{
	return len(a)
}
func (a ByPrice) Swap(i, j int){
	a[i], a[j] = a[j], a[i]
}
func (a ByPrice) Less(i, j int) bool {
	return a[i].Price < a[j].Price
}

func main(){
	mobile := []Mobile{
		{"Sony", 952},
		{"Nokia", 468},
		{"Apple", 1219},
		{"Samsung", 1045},
	}
	fmt.Println("before")
	for _, v:=range mobile{
		fmt.Println(v.Brand, v.Price)
	}
	sort.Sort(ByPrice(mobile))
	for _, v:=range mobile{
		fmt.Println(v.Brand, v.Price)
	}
}