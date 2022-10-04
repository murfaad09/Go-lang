package main

import (
	"fmt"
	"sync"
)

var G = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	G = G + 1
	m.Unlock()
	wg.Done()
}
func main() {
	var mg sync.Mutex
	var w sync.WaitGroup
	for i := 0; i < 100; i++ {
		w.Add(1)
		worker(&w, &mg)
	}
	w.Wait()
	fmt.Println("Value of x:", G)
}
