package main

import (
	"encoding/json"
	"fmt"
)

type M  map[string]interface{}

func main(){
	currency := map[string]int{
		"one":1,
		"two":2,
	}
	countrydetails := map[string]map[string]int{
		"bangladesh":{"taka":1},
		"america":{"usd":5},
	}
	fmt.Println(currency)
	fmt.Println("")
	fmt.Println(countrydetails)

	var myMapslice []M
	m1 := M{"dn":"abc","status":"live"}
	m2 := M{"dn":"version","status":"inactive"}
	myMapslice = append(myMapslice, m1, m2)
	myJson, _ := json.MarshalIndent(myMapslice,""," ")
	fmt.Println(string(myJson),)
	fmt.Println([]string{"one","two"})
}

// ["one","two","three"]
// [[0 1 2 3] [0 1 2 3] [0 1 2 3]]
// {"one","two"}
// {"one":5,"two":2} key, value
// [] [{"one":1,"two":2},{"three":3,"four":4}] key, value - list

//  {
//	"Great Britain Pound": {"GBP": 1},
//	"Euro":                {"EUR": 2},
//	"USA Dollar":          {"USD": 3},
//	}


//  [{
//	"Great Britain Pound": {"GBP": 1},
//	"Euro":                {"EUR": 2},
//	"USA Dollar":          {"USD": 3},
//	},
//  {
//	"Great Britain Pound": {"GBP": 1},
//	"Euro":                {"EUR": 2},
//	"USA Dollar":          {"USD": 3},
//	}]

// add
// delete
// update
// pop
// search
// sort

// set of instruction that computer execute

// data type
// variables
// keywords
// logical and arithmetical operations
// if else conditions
// loops
// number , character, arrays
// functions
// input and output operations
// class