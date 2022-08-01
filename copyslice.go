package main

import "fmt"

func main() {
	mainslice := [10]string{"Ali", "Hamza", "Usama", "Haseeb", "Basit", "Hamza", "Faisal", "Bilal", "Usman", "Ayaz"}

	slice2 := mainslice[:len(mainslice)-5]
	//slice3 := make([]string, len(slice2))
	//copy(slice3, slice2)
	slice3 := slice2
	fmt.Printf("Length: %d", len(slice3))
	fmt.Println("")
	fmt.Printf("Value: %v", slice3)
	fmt.Println("")
	fmt.Printf("Capacity: %d", cap(slice3))
}
