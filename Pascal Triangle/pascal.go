package main

import (
	"fmt"
)

// arr := []int{}
func main() {
	inputSlice := []int{2, 0, 2, 1, 1, 0}
	sliceZ := make([]int, 0)
	slice1 := make([]int, 0)
	slice2 := make([]int, 0)
	for i := 0; i < 6; i++ {
		if inputSlice[i] == 0 {
			sliceZ = append(sliceZ, inputSlice[i])

		}
		if inputSlice[i] == 1 {
			slice1 = append(slice1, inputSlice[i])
		}
		if inputSlice[i] == 2 {
			slice2 = append(slice2, inputSlice[i])
		}

	}
	sliceF := make([]int, 0)
	for _, val := range sliceZ {
		sliceF = append(sliceF, val)
	}
	for _, val := range slice1 {
		sliceF = append(sliceF, val)
	}
	for _, val := range slice2 {
		sliceF = append(sliceF, val)
	}
	fmt.Println(sliceF)
}
