package main

import "fmt"

type MyTestInterface interface {
	TotalPrice() float64
}

type Forest struct {
	Name string
	Area float64
	TreeList []string
}

type Amazon struct {
	Forest
	area bool
}
func main(){
	ape := []string{"monkey","Eagle","Lucky"}
	tiger := [][]string{ape,ape}
	george := map[string]string{
		"ape":"name george",
		"age": "may be die",
	}
	fmt.Println(tiger)
	fmt.Println(george)
	rampage := map[string]map[string]string{
		"one": {"name":"animal","age":"500"},
		"two": {"name":"bird", "age":"300"},
	}
	helicopter := map[string][]string{
		"gOne" : []string{"helicoptor one","helicoptor two"},
		"gTwo" : []string{"car one","car two"},
	}
	fmt.Println(rampage)
	fmt.Println(helicopter)
	viot := struct {
		Name string
		characteris []string
	}{
		Name:"nayon",
		characteris:[]string{"software engineer","programmer"},
	}
	fmt.Println(viot)

	var employee = make(map[string]int)
	employee["one"] = 1
	employee["two"] = 2
	fmt.Println(employee)
	delete(employee,"one")
	fmt.Println(employee)
	func(l int, b int){
		fmt.Println(l*b)
	}(20,30)

}
func squareSum(x int) Second{
	return func(y int) First{
		return func(z int) int{
			return x*x + y*y
		}
	}
}
func rectangle(l int, b int)(area int){
	var parameter int
	parameter = 2*(l+b)
	fmt.Println(parameter)
	area = l*b
	return
}
type First func(int) int
type Second func(int) First
var(
	area = func(l int, b int) int{
		return l*b
	}
)

func rectangle1(l int, b int)(area int, parameter int){
	parameter = 2 * 3
	area = l * b
	return
}