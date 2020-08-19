package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int32
	wg sync.WaitGroup
)

func increment(name string){
	defer wg.Done()
	for range name {
		atomic.AddInt32(&counter, 1)
		runtime.Gosched()
	}
}

func main(){
	wg.Add(3)
	go increment("python")
	go increment("java")
	go increment("golang")
	wg.Wait()
	fmt.Println("counter", counter)
}