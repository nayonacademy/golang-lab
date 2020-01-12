package main

import (
	"fmt"
	"runtime"
	"time"
)

var i, j int = 1, 2
var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
)

func add(x int, y int) int {
	return x + y
}

func main() {
	cards := []string{"one", "two"}
	cards = append(cards, "three")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}

	var i2 int
	var f float32
	var s string
	var bl bool
	var c, b, a = true, false, "no"
	var i, j int = 1, 2
	k := 3
	var z float32 = float32(j)
	v := 42
	d, python, java := true, false, "no0"
	const Pi = 3.14
	fmt.Println(add(3, 5))
	fmt.Println(swap("hello", "world"))
	fmt.Println(split(14))
	fmt.Println(c, b, a)
	fmt.Println(d, python, java, i, j, k)
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("type: %T and Value : %v\n", ToBe, ToBe)
	fmt.Printf("%v %v %q %v", i2, f, s, bl)
	fmt.Printf("Conversion %T = %v\n", z, z)
	fmt.Printf("v is of type %T\n", v)
	fmt.Println("Happy ", Pi, " Day")

	sum := 0
	for i := 0; i < 10; i++ {
		sum += 1
		fmt.Println(sum)
	}

	sm := 1
	for sm < 9 {
		sm += sm
	}
	fmt.Println('\n', sm)

	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)

	}

	today := time.Now().Weekday()
	fmt.Println(today)
	vtx := Vertex{4, 5}
	fmt.Println(vtx.X)

	var ab [2]string
	ab[0] = "Hello"
	ab[1] = "World"
	fmt.Println(ab[0], ab[1])
	primes := [7]int{2, 4, 6, 9, 23, 324, 23}
	var sss []int = primes[1:3]
	if sss == nil {
		fmt.Println("its null")
	} else {
		fmt.Println("It has data")
	}
	fmt.Println(sss)
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 5
	y = sum - x
	return
}

func pow(x, n, lim float64) float64 {
	if v := 4.00 + 3.00; v < lim {
		return v
	}
	return lim
}

type Vertex struct {
	X int
	Y int
}
