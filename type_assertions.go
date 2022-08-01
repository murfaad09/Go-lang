package main

import "fmt"

func main() {
	var l interface{} = "hellow"
	z := l.(string)
	fmt.Println(z)

	z, ok := l.(string)
	fmt.Println(z, ok)
}
