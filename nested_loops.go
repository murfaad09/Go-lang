package main

import "fmt"

func main() {
	//Nested loops is loop inside a loop
	var array = [3]string{"red", "yellow", "white"}
	var array2 = [3]string{"Apple", "Berry", "Mango"}
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array2); j++ {
			fmt.Println(array[i], "", array2[j])
		}
	}
}
