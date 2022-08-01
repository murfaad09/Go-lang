package main

import "fmt"

func main() {
	var array = []int{1, 25, 13, 94, 55, 16, 677, 68, 39, 10}
	slice := array[3:8]
	fmt.Println(slice)
}
