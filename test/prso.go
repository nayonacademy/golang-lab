package main

import "fmt"

type Person struct {
	firstname string
	lastname  string
	age       int
}

func (p *Person) updateName(Newname string) {
	(*p).firstname = Newname
}
func (p Person) print() {
	fmt.Println(p)
}