package main

import "fmt"

type Creature struct {
	Name string
	Real bool
}

func Dump(c*Creature){
	fmt.Printf("Name %s, Real: %t\n", c.Name, c.Real)
}

func (c Creature) Dump(){
	fmt.Printf("Name %s, Real: %t\n", c.Name, c.Real)
}

type FlyingCreature struct {
	Creature
	WingSpan int
}

func (fc FlyingCreature) Dump(){
	fmt.Printf("Name %s, Real: %t Wingspan: %d\n",fc.Name, fc.Real, fc.WingSpan )
}

type Unicorn struct {
	Creature
}

type Dragon struct {
	FlyingCreature
}

type Pterodactyl struct {
	FlyingCreature
}

func NewPterodactyl(wingspan int) *Pterodactyl  {
	pet := &Pterodactyl{FlyingCreature{
		Creature: Creature{"Pterodactyl",true},
		WingSpan: 0,
	},}
	return pet
}

type Dumper interface {
	Dump()
}
type Door struct {
	Thickness int
	Color string
}

func (d Door) Dump(){
	fmt.Printf("Door Thickness: %d, Color: %s", d.Thickness, d.Color)
}

func main(){
	creature := &Creature{
		Name: "some creature",
		Real: false,
	}
	uni := Unicorn{Creature{
		Name: "Unicorn",
		Real: false,
	}}
	pet1 := &Pterodactyl{FlyingCreature{
		Creature: Creature{"pterodactyl",true},
		WingSpan: 5,
	}}
	pet2 := NewPterodactyl(8)
	door := &Door{
		Thickness: 3,
		Color:     "red",
	}
	Dump(creature)
	creature.Dump()
	uni.Dump()
	pet1.Dump()
	pet2.Dump()
	creatures := []Creature{
		*creature,
		uni.Creature,
		pet1.Creature,
		pet2.Creature,
	}
	fmt.Println("Dump() through creature embeded type")
	for _, creature := range creatures{
		creature.Dump()
	}
	dumpers := []Dumper{creature, uni, pet1, pet2, door}
	fmt.Println("dump through interface")
	for _, dumper := range dumpers{
		dumper.Dump()
	}
}