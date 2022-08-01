package main

import "fmt"

func main() {
	slice := []int{22, 34, 55, 45, 32}
	slice = append(slice, 888, 978)
	fmt.Printf("Length:%d", len(slice))
	fmt.Println("")
	fmt.Printf("Slice value:%v", slice)
	fmt.Println("")
	fmt.Printf("Capacity:%v", cap(slice))
	fmt.Println("")
}
