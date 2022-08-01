package main

import "fmt"

func main() {
	slice := make([]int, 6, 10)
	fmt.Printf("Type:%T", slice)
	fmt.Println("")
	fmt.Printf("Length:%d", len(slice))
	fmt.Println("")
	fmt.Printf("Capacity:%d", cap(slice))
}
