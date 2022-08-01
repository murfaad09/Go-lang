package main

import "fmt"

//function with two parameters
func demo(x int, y string) {
	fmt.Println("Name:", y, "Id:", x)
}

func main() {
	// calling a function
	demo(1234, "InvoZone")
	demo(2524, "Golang")
}
