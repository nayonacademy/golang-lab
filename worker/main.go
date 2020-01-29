package main

import "fmt"

import "time"

func worker(id int, jobs chan int, results chan int) {
	for j := range jobs {
		fmt.Println("Worker id : ", id, " start")
		time.Sleep(time.Second)
		fmt.Println("Worker id ", id, " finish")
		results <- j * 2
	}
}

func main() {
	const numsjob = 5
	jobs := make(chan int, numsjob)
	results := make(chan int, numsjob)
	for w := 1; w < 3; w++ {
		go worker(w, jobs, results)
	}
	for j := 1; j < numsjob; j++ {
		jobs <- j
	}
	close(jobs)
	for a := 1; a < numsjob; a++ {
		<-results
	}
}

// https://gobyexample.com/worker-pools
