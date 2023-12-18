package main

import "fmt"

func maain() {
	c := make(chan int, 3)

	c <- 1
	c <- 2
	c <- 3

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println(<-c)
}
