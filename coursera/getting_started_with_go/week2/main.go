package main

import "fmt"
type Grade int
const (
	A Grade = iota
	B
	C
	D
	F
)


func main(){
	var x int32 = 1
	var y int16 = 2
	fmt.Println(x,y)
	if x > 5 {
		fmt.Println("Yup")
	}

	for i:=0;i<10;i++{
		fmt.Println("Hi ")
	}
	a = 0
	for i < 10{
		fmt.Println("hi ")
		a++
	}
	for {
		fmt.Println("hi ")
	}

	switch c {
	case 1:
		fmt.Println("Case 1")
	case 2:
		fmt.Println("Case 2")
	case 3:
		fmt.Println("Case 3")
	}

	switch {
	case x > 1:
		fmt.Println("Case 1")
	case x < -1:
		fmt.Println("case 2")
	}

	e := 0
	for i< 10{
		e++
		if e == 5{
			break
		}
		fmt.Println("hi ")
	}

}
