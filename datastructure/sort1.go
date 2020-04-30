package main

import (
	"fmt"
	"sort"
)

func main(){
	intSlide := []int{12,34,11,7,34,87}
	fmt.Println(intSlide)
	sort.Ints(intSlide)
	fmt.Println(intSlide)
	strSlide := []string{"ox", "banana","apple","yoga"}
	fmt.Println(strSlide)
	sort.Strings(strSlide)
	fmt.Println(strSlide)
	fmt.Println(sort.IntsAreSorted(intSlide))
	fmt.Println(sort.StringsAreSorted(strSlide))
	x := 34
	fmt.Println(sort.SearchInts(intSlide, x))
}