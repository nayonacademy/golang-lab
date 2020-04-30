package main

import (
	"fmt"
	"sort"
)

func main(){
	a := []int{15, 4, 33, 52, 551, 90, 8, 16, 15, 105}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	a = []int{-15, -4, -33, -52, -551, -90, -8, -16, -15, -105}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a)
	b := []string{"Montana","Alaska","Indiana","Nevada","Washington","Ohio","Texas"}
	sort.Sort(sort.Reverse(sort.StringSlice(b)))
	fmt.Println(b)
	s := []float64{85.201, 14.74, 965.25, 125.32, 63.14}
	sort.Sort(sort.Float64Slice(s))
	fmt.Println(s)
	fmt.Println(sort.Float64Slice(s).Search(85.201))
}
