package main

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"
)

type Envelope struct {
	Type string
	Msg interface{}
}

type Sound struct {
	Description string
	Authority string
}

type Cowbell struct {
	More bool
}
const input = `
{
	"type": "sound",
	"msg": {
		"description": "dynamite",
		"authority": "the Bruce Dickinson"
	}
}`

func main(){
	sammy := map[string]string{"name":"nayon","color":"light","location":"dhaka"}
	items := make([]string, len(sammy))

	var i int
	for _, v := range sammy{
		items[i] = v
		i++
	}
	itm := []string{"two","four","eight","one"}
	fmt.Printf("%q", items)
	sort.Strings(itm)
	fmt.Println(itm)
	fmt.Println(strings.Repeat("#",5))
	s := Envelope{
		Type: "sound",
		Msg:  Sound{
			Description: "dynamic",
			Authority:   "the bruce dickinson",
		},
	}
	buff, err := json.Marshal(s)
	if err != nil{
		log.Fatal(err)
	}
	c := Envelope{
		Type: "cowbell",
		Msg:  Cowbell{More:true},
	}
	fmt.Printf("%s\n", buff)
	fmt.Printf("%s", c)

	var env Envelope
	if err := json.Unmarshal([]byte(input), &env); err != nil{
		log.Fatal(err)
	}
}