package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {

	// var diamond = []string{"diamond_2", "diamond_3", "diamond_4", "diamond_5", "diamond_6", "diamond_7", "diamond_8", "diamond_9", "diamond_10", "diamond_J", "diamond_Q", "diamond_K", "diamond_A"}
	// for i := 0; i < len(diamond); i++ {
	// 	fmt.Println("Diamond value:", diamond[i])
	// }

	diamond := []string{"heart_2", "heart_3", "heart_4", "heart_5", "heart_6", "heart_7", "heart_8", "heart_9", "heart_J", "heart_Q", "heart_K", "heart_A", "diamond_2", "diamond_3", "diamond_4", "diamond_5", "diamond_6", "diamond_7", "diamond_8", "diamond_9", "diamond_10", "diamond_J", "diamond_Q", "diamond_K", "diamond_A"}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(diamond), func(i, j int) {
		diamond[i], diamond[j] = diamond[j], diamond[i]
	})
	fmt.Println("After shuffling diamond:", diamond)
	fmt.Println("")
	sort.Slice(diamond, func(p, q int) bool {
		return diamond[p] < diamond[q]
	})

	fmt.Println("Sort diamond according to their names:")
	fmt.Println(diamond)
}
