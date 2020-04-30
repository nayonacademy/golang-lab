package main

import "fmt"

type Fruits struct {
	Name string
	Amount float64
}

type Flowers struct {
	Name string
	Amount float64
}

type MyInterface interface {
	Price() float64
}

func (f Fruits) Price() float64{
	return f.Amount * 19
}

func main(){
	f := Fruits{
		Name:   "apple",
		Amount: 20,
	}
	fmt.Println(f)
	fmt.Println(f.Amount)

	bio := struct {
		Name string
		Friends map[string]int
		Foods []string
	}{
		Name:"Nayon",
		Friends: map[string]int{
			"a":123213,
			"b":24234,
		},
		Foods:[]string{
				"mango",
				"banana",
		},
	}
	fmt.Println(bio.Name)
	earth := SolarSystem{
		Planet: Planet{
			Name:"Earth",
			Distance:123434.00,
			Character: []string{"mother earth","Lovely"},
		},
		status: false,
	}
	fmt.Println(earth)
	tshirt := configurable{}
	tshirt.price = 1550
	tshirt.qty = 2
	fmt.Println(tshirt.shipping())
	fmt.Println(tshirt.tax())
	fmt.Println(tshirt.offer())
	fmt.Println(tshirt.available())
}

type Planet struct {
	Name string
	Distance float64
	Character []string
}

type SolarSystem struct {
	Planet
	status bool
}

type discount interface {
	offer() float64
}
type giftpack interface {
	available() string
}
type catalog interface {
	discount
	giftpack
	shipping() float64
	tax() float64
}
type configurable struct {
	name string
	price, qty float64
}

func (c *configurable) tax() float64  {
	return c.price * c.qty
}
func (c *configurable) shipping() float64{
	return c.qty *5
}

func (c *configurable) offer() float64  {
	return c.price * 0.15
}

func (c *configurable) available() string{
	if c.price > 1000{
		return "available"
	}
	return "not available"
}
type download struct {
	name string
	price, qty float64
}
func (d *download) tax() float64{
	return d.price *d.qty
}
func (d *download) available() string{
	if d.price > 500{
		return "available"
	}
	return "not available"
}
type simple struct {
	name string
	price, qty float64
}
func (s *simple) tax() float64{
	return s.price * s.qty
}

func (s *simple) shipping() float64{
	return s.qty * 3
}
func (s *simple) offer() float64{
	return s.price * .10
}