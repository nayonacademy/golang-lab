package main

import "fmt"

func main(){
	integerArray := []int{10,20,30,40,50}
	stringArray := [3]string{"fitst","second","third"}
	fmt.Println(integerArray, stringArray)
	for key, value := range integerArray{
		fmt.Println(key, value)
	}
	for _, val := range stringArray{
		fmt.Println(val)
	}
	for i:=0; i < len(stringArray); i++{
		fmt.Println(stringArray[i])
	}
	nested := make([][]int, 0, 3)
	for n:=0;n<3; n++{
		out := make([]int, 0, 4)
		for j:=0; j<4;j++{
			out = append(out, j)
		}
		nested = append(nested, out)
	}
	fmt.Println(nested)
//	Example 2
	appleLaptops := []string{"macbookPro","macbookair"}
	hplaptopr := []string{"hp","hpelite"}
	laptosp := [][]string{appleLaptops, hplaptopr}
	fmt.Println(laptosp)
	for a,b := range laptosp{
		fmt.Println("Records : ", a)
		for _, name := range b{
			fmt.Printf("\t laptop name: %v \n", name)
		}
	}
	nameAgeMap := map[string]int{
		"James":30,
		"Ali": 39,
	}
	fmt.Println(nameAgeMap)
	for key, value := range nameAgeMap{
		fmt.Printf("%v is %d years old\n", key, value)
	}
	currencey := map[string]string{
		"aud":"australia",
		"gbp" : "great britain",
	}
	currencey["usd"] = "USD dollar"
	fmt.Println(currencey)
	delete(currencey,"gbp")
	fmt.Println(currencey)
	currencey["aud"] = "bangladesh"
	fmt.Println(currencey)
	for key, value := range currencey{
		fmt.Printf("%v might be eqal %v\n", key, value)
	}
	if _, ok := currencey["usd"]; ok{
		fmt.Println(ok)
	}else{
		fmt.Println("we dont find the value")
	}
	nameAndHobby := map[string][]string{
		"one": []string{"two","three","four"},
		"fruits": []string{"apple","orange","jackfruits"},
	}
	fmt.Println(nameAndHobby)
	nameAndHobby["timit"] = []string{"cartoon","laughing"}
	fmt.Println(nameAndHobby)
	for i, v := range nameAndHobby{
		fmt.Printf("%v likes \n",i)
		for j, val := range v{
			fmt.Printf("\t%v %v\n",j, val)
		}
	}
	money := map[string]map[string]int{
		"Britain" : {"GBP":1},
		"Euro": {"EUR":2},
		"Dollar": {"USD":3},
	}
	for key, value := range money{
		fmt.Printf("Currency name : %v\n", key)
		for k, v := range value{
			fmt.Printf("\t Currency Code: %v \n\t\t Ranking: %v\n\n", k, v)
		}
	}
	animal1 := Animal{
		name:            "lion",
		characteristics: []string{"wild","jungle"},
	}
	fmt.Println("Animal Name: ", animal1.name)
	for _, v:=range animal1.characteristics{
		fmt.Printf("\t %v\n", v)
	}
	herb := Herbivore{Animal:Animal{
		name:            "Goat",
		characteristics: []string{"lazy","grass"},
	},
	eatHuman:false,
	}
	fmt.Println(herb.Animal.characteristics)
	fmt.Println(herb.eatHuman)
//
	bio := struct {
		firstName string
		friends map[string]int
		favDrinks []string
	}{
		firstName: "steven",
		friends: map[string]int{
			"tim": 123233,
			"abdul": 23434,
		},
		favDrinks:[]string{
				"Pepsi",
				"Cocacola",
		},
	}
	fmt.Println(bio.firstName)
	for k, v := range bio.favDrinks{
		fmt.Println(k,v)
	}
	for k, v := range bio.friends{
		fmt.Println(k,v)
	}
}

type Animal struct {
	name string
	characteristics []string
}

type Herbivore struct {
	Animal
	eatHuman bool
}