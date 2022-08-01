package main

import "fmt"

func main() {
	cars := []string{"Mercedez", "Audi", "Lexus"}
	for idx, val := range cars {
		fmt.Printf("%v\t%v\n", idx, val)
	}
}
