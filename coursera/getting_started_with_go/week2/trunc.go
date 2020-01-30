package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	fmt.Scan(&x)
	x = math.Trunc(x)
	fmt.Println(x)
}