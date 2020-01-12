package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}
type person struct {
	firstname string
	lastname  string
	contract  contactInfo
}

func (d person) print() {
	fmt.Printf("%+v", d)
}

func (pointerToPerson *person) updateName(newFirstname string) {
	(*pointerToPerson).firstname = newFirstname
}
func main() {
	mySlice := []string{"hi", "how", "are"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
	jim := person{
		firstname: "jim",
		lastname:  "party",
		contract: contactInfo{
			email:   "jim@j.com",
			zipcode: 9000,
		},
	}
	jimPointer := &jim

	jimPointer.updateName("jimmy")
	jim.print()
}

func updateSlice(s []string) {
	s[0] = "bye"
}
