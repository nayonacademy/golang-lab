package main

import (
	"fmt"
)

type Fruits map[string]interface{}

type Thing struct {
	name string
	age int
}

func main(){
	//garden := []Fruits{
	//	{
	//		"one": []Fruits{
	//			{"apple":1},{"banana":2},
	//		},
	//	},
	//	{
	//		"one": []Fruits{
	//			{"apple":3},{"banana":4},
	//		},
	//	},
	//}
	////data, _ := json.Marshal(garden)
	////fmt.Printf("%s",data)
	//for i := range garden{
	//	fmt.Println(json.Marshal(garden[i]))
	//}
	//
	//x := make(map[string][]string)
	//
	//x["key"] = append(x["key"], "value")
	//x["key"] = append(x["key"], "value1")
	//
	//fmt.Println(x)
	//fmt.Println(x)

	myname := map[int][]Thing{}

	for i := 0; i < 5; i++ {
		myname[i] = append(myname[i], Thing{
			name: "apple",
			age:  24+i,
		})
	}

	//fmt.Println(myname)
	for i,v := range myname{
		fmt.Println(i, " ",v)
	}

	s := []int{1,2,3}
	for i := range s{
		s[i] += 1
	}
	fmt.Println(s)

	var es = []*A{
		&A{
			a:0,
			s:"test",
		},
		&A{
			a:1,
			s:"test1",
		},
	}
	for _, e := range es{
		fmt.Printf("%v", e)
		e.a++
	}
	for _, e := range es{
		fmt.Printf("%v", e)
		e.a++
	}
	for i, j := 0,1;i<10;i,j = i+1,j+1{
		fmt.Println("hello, playground", i, j)
	}
	sammyShark := map[string]string{"name":"Sammy","animal":"shark","color":"red"}
	for key, val := range sammyShark{
		fmt.Println(key + ": " + val)
	}
	for i:=3; i < 17; i +=4 {
		fmt.Println(i)
	}
	for i:= 100; i>1; i -= 15{
		fmt.Println(i)
	}
	i := 0
	for i < 5{
		fmt.Println(i)
		i++
	}
	sharks := []string{"one","apple","orange","banana"}
	for _, shar := range sharks{
		fmt.Println(shar)
	}
	for range sharks{
		sharks = append(sharks, "shark")
	}
	fmt.Printf("%q\n", sharks)

	integers := make([]int, 10)
	fmt.Println(integers)
	for i := range integers{
		integers[i] = i
	}
	fmt.Println(integers)
	sammy := "Sammy"
	for _, letter := range sammy{
		fmt.Printf("%c\n", letter)
	}
	fmt.Println("\n")

	numList := []int{1,2,3}
	alphaList := []string{"a","b","c"}
	for _, i := range numList{
		fmt.Println(i, "#####")
		for _, letter := range alphaList{
			fmt.Println(letter)
		}
	}
	ints := [][]int{
		[]int{0,1,2,3},
		[]int{-1,-2,-3},
		[]int{9,8,7},
	}
	for _, i := range ints{
		fmt.Println(i)
	}

}

type A struct {
	a int
	s string
}