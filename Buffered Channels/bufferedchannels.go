package main

import "fmt"

func main() {
	c := 4
	ch := make(chan int, c)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	for i := 1; i <= c; i++ {
		fmt.Println(<-ch)
	}
}
