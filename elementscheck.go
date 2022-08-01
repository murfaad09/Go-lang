package main

import "fmt"

func main() {
	var a = map[int]int{1: 100, 2: 200, 3: 300, 4: 400}
	val1, ok1 := a[1]
	val2, ok2 := a[5]
	fmt.Println(val1, ok1)
	fmt.Println(val2, ok2)
}
