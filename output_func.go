package main

import "fmt"

func main() {
	var print string = "Hello"
	var printi string = "World"
	fmt.Print("The value of variable will be show like this with simple print:>", print, printi)
	fmt.Println("The value of variable will be show like this with println:>", print, printi)
	fmt.Printf("The value of variable will be show like this with printf:> value: %v...type:%T\n", print, printi)

}
