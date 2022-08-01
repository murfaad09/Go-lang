package main

import "fmt"

func ts(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Type is int:value", v)
	case string:
		fmt.Println("Type is string", v)
	default:
		fmt.Println("Type is neither int nor string")
	}

}
func main() {
	ts(23)
	ts("Hiii")
}
