package main

import "fmt"

func main() {
	//Break statement terminates the program execution when a specific statement occurs
	for i := 0; i < 10; i++ {
		if i == 6 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("")
	fmt.Println("Program terminated when i=6")
}
