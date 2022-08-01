package main

import "fmt"

func main() {
	var size int = 22
	//fmt.Println("Loop is used for repeatition of some task again and again untill we get desired results.")
	//fmt.Println("This loop will excute it self and controlled by a controll variable")
	for i := 0; i < size; i++ {
		fmt.Printf("VAlue of i= %d  and in iteration i= %d\n", i, i)
	}
}
