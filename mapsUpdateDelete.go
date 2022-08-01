package main

import "fmt"

func main() {
	var a = make(map[int]string)
	a[1] = "Helloo"
	a[2] = "Hi"
	a[3] = "Hiiii"

	for i := 1; i < 4; i++ {
		fmt.Println("value", a[i])
	}
	//updating
	a[1] = "SUnday"
	fmt.Print(a[1])

	// Delete
	delete(a, 2)
	fmt.Println("")
	fmt.Println(a)
}
