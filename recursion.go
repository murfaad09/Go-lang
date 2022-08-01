package main

import "fmt"

func recursion(z int) int {
	if z == 10 {
		return 0
	}
	fmt.Println(z)
	return recursion(z + 1)
}
func main() {
	recursion(2)
}
