package main

import (
	"fmt"
)

func main() {
	name := [4][2]string{{"ali", "Usman"}, {"Ahmed", "Bilal"}, {"Arslan,Hamza"}, {"Basit", "Arslan"}}
	xindex := make([]string, len(name))
	yindex := make([]string, len(name[0]))
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
	for i := 0; i < len(name); i++ {
		for j := 0; j < len(name[i]); j++ {
			if name[i][j] == "ali" {
				//name[i][j] = "hello"
				xindex[i] = "hello"
				yindex[j] = "hello"

			}

		}

	}
	fmt.Println("After rearranging the matrix")
	for i := 0; i < len(name); i++ {
		for j := 0; j < (len(name[i])); j++ {
			if xindex[i] == "hello" {
				name[i][j] = "Hello"

				//xindex[i] = "Hi"
				//yindex[j] = "Hi"
			} else if yindex[j] == "hello" {
				name[i][j] = "Hello"

			}
		}

	}
	for i := 0; i < len(name); i++ {
		fmt.Println(name[i])
	}
}
