package main

import "fmt"

func Worker(id int, jobs <-chan int, results chan int) {
	for job := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", id, job)
		fib := Fibonacci(job)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", id, job, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func maaaain() {
	task := []int{2, 3, 4, 5, 7, 19, 12, 40}
	nWorkers := 3
	jobs := make(chan int, len(task))
	results := make(chan int, len(task))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, value := range task {
		jobs <- value
	}
	close(jobs)

	for r := 0; r < len(task); r++ {
		<-results
	}
}
