package main

import (
	"fmt"
	"net/http"
)

//type geometry interface {
//	area() float64
//	perim() float64
//}
//
//type react struct {
//	width, height float64
//}
//
//type circle struct {
//	radius float64
//}
//
//func (r react) area() float64{
//	return r.height *r.width
//}
//
//func (r react) perim() float64{
//	return 2*r.width + 2*r.height
//}
//
//func (c circle) area() float64{
//	return math.Pi * c.radius * c.radius
//}
//
//func (c circle) perim() float64{
//	return 2 * math.Pi * c.radius
//}

//func measure(g geometry){
//	fmt.Println(g)
//	fmt.Println(g.area())
//	fmt.Println(g.perim())
//}
//
//func main(){
//	r := react{width:3, height:5}
//	c := circle{radius:5}
//	measure(r)
//	measure(c)
//}

//

type todo int
func (d todo)ServeHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "you want to show somethings")
}

func main(){
	var d todo
	http.ListenAndServe(":8080",d)
}