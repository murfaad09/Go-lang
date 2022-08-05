package main

import "fmt"

func fab(u int, c chan int) {
	x, y := 0, 1
	for i := 0; i < u; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
func main() {
	c := make(chan int, 10)
	go fab(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
