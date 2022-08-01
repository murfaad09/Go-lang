package main

import "fmt"

func mularray(x int, y string) (integer int, text string) {
	fmt.Println("Integer Value and String value is displayed below")
	integer = x + x
	text = y + "World"
	return
}

func main() {

	fmt.Println(mularray(5, "Hello"))
}
