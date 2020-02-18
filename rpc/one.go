package main

import "fmt"

type ItemList struct {
	ID string
	Name string
}
var storage []ItemList
func AddItemList(name ItemList) ItemList{
	storage = append(storage, name)
	return name
}
func GetbyId(idno string)ItemList{
	var items ItemList
	for _, id := range storage{
		if id.ID == idno{
			items = id
		}
	}
	return items
}
func ShowList(){
	fmt.Println(storage)
}

func main(){
	a := ItemList{ID:"1", Name:"Hello"}
	b := ItemList{ID:"2", Name:"Hello second"}
	AddItemList(a)
	AddItemList(b)
	ShowList()
	c := GetbyId("2")
	fmt.Println(c)
}