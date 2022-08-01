package main

import "fmt"

type Car struct {
	name    string
	model   int
	company string
}

func main() {

	var variable1 Car
	variable1.name = "Mustang X"
	variable1.model = 2022
	variable1.company = "Mustang"

	var variable2 Car
	variable2.name = "BMW"
	variable2.model = 2021
	variable2.company = "BMW"

	fmt.Println("Name:", variable1.name)
	fmt.Println("Model:", variable1.model)
	fmt.Println("Company:", variable1.company)
	fmt.Println("")
	fmt.Println("Name:", variable2.name)
	fmt.Println("Model:", variable2.model)
	fmt.Println("Company:", variable2.company)

}
