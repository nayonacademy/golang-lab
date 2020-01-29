package main

import "fmt"

// Person struct
type Person struct {
	firstName string
	lastName  string
	age       int32
}

func (p Person) print() {
	fmt.Printf("+%v", p)
}

// UpdateName use
func (p *Person) UpdateName(fname string) {
	(*p).firstName = fname
}
func main() {
	jim := Person{
		firstName: "apple",
		lastName:  "ornage",
	}
	jim.UpdateName("banana")
	jim.print()
}
