package main

import (
	"encoding/json"
	"fmt"
	"sort"
)

type Jsotest struct {
	Id string `json:"id"`
	Name string `json:"name"`
}

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

	name := make(map[string]interface{})
	name["id"] = map[string]int{"t":2,"r":2}
	name["name"] = "two"

	fmt.Println(name)

//	convert map to json
	jsonString, _ := json.Marshal(name)
	fmt.Println(string(jsonString))

//	convert json to map
	s := Jsotest{}
	json.Unmarshal(jsonString, &s)
	fmt.Println(s)
}