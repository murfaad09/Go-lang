package main

import "fmt"

func main() {
	price := []int{2, 4, 8, 10, 45, 60, 1, 2}

	var buy int
	var sell int
	var profit int = 0

	for i := 0; i < 8; i++ {

		for j := i; j < 8; j++ {

			buy = price[i]
			sell = price[j]
			if profit < sell-buy {

				profit = sell - buy
			}

		}

	}
	fmt.Println("value", profit)

}
