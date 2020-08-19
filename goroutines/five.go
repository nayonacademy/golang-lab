package main

import (
	"fmt"
	"sync"
)

var (
	counter1 int32
	wg1 sync.WaitGroup
	mutex sync.Mutex
)

func increamentOne(name string){
	defer wg1.Done()
	for i := 0; i<5; i++{
		mutex.Lock()
		{
			counter1++
		}
		mutex.Unlock()
	}
}

func main()  {
	wg1.Add(3)
	go increamentOne("python")
	go increamentOne("golang")
	go increamentOne("java")
	wg1.Wait()
	fmt.Println(counter1)
}
