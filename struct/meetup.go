package main

import "fmt"
type PersonNew interface {
	SayHello()
	GetDetails()
}
type PersonNew struct {
	name string
	age int
	city, phone string
}

func (p PersonNew) SayHello(){
	fmt.Println("Hello %s", p.name)
}

func (p PersonNew) GetDetails(){
	fmt.Println("Hello Mr. %s you live in %s and your phone no is %s , now you are %d age old",p.name, p.city, p.phone,p.age)
}
type Speaker struct {
	person PersonNew
	speaksOn []string
	pastEvents []string
}
func (s Speaker) GetDetails(){
	s.person.GetDetails()
	fmt.Println("Speaker talks on following technologies:")
	for _, value := range s.speaksOn {
		fmt.Println(value)
	}
	fmt.Println("Presented on the following conferences:")
	for _, value := range s.pastEvents {
		fmt.Println(value)
	}
}
type Organization struct {
	PersonNew
	meetups []string
}
func (o Organization) GetDetails(){
	o.PersonNew.GetDetails()
	for _, value := range o.meetups{
		fmt.Println(value)
	}
}
func main(){
	//p := PersonNew{name:"newton",age:60, city:"London", phone:"29864982374"}
	//fmt.Println(p)
}