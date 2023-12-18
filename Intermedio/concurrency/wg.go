package main

import (
	"fmt"
	"sync"
	"time"
)

func maaain() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go doSomethingg(i, &wg)
	}

	wg.Wait()
}

func doSomethingg(i int, wq *sync.WaitGroup) {
	defer wq.Done()
	fmt.Println("Startd", i)
	time.Sleep(2 * time.Second)
	fmt.Println("Finish", i)
}
