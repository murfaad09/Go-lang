package main

import "fmt"

func channels(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	c := make(chan int)
	go channels(s[:len(s)/2], c)
	go channels(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x + y)

}
