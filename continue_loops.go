package main

import "fmt"

func main() {
	//Continue statement is used to skip one or more statements

	var size int = 12
	for i := 0; i < size; i++ {
		if i == 10 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("")
	fmt.Println("10 is skipped")
}
