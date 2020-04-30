package main

import (
	"fmt"
)

type gadgets uint8
const(
	Camera gadgets = iota
	Bluetooth
	Media
	Storage
	VideoCalling
	Multitasking
	Messaging
)

type mobile struct {
	make string
	model string
}

type smartphone struct {
	gadgets gadgets
}

func (s *smartphone) launch(){
	fmt.Println("New smartphone")
}
type android struct {
	mobile
	smartphone
	waterproof string
}
func (a *android) samsung(){
	fmt.Print("%s %s\n", a.make, a.model)
}

type iphone struct {
	mobile
	smartphone
	sensor int
}
func (i *iphone) apple(){
	fmt.Printf("%s %s\n", i.make, i.model)
}

func main(){
	t := &android{}
	t.make = "samsung"
	t.model = "j7"
	t.gadgets = Camera+Bluetooth+Media+Storage+VideoCalling+Multitasking
	t.launch()
	t.samsung()
}