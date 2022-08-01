package main

import "fmt"

type I interface {
	ff()
}
type T struct {
	x string
}

func (t T) ff() {
	fmt.Println(t.x)
}
func main() {
	var i I = T{"Hello"}
	i.ff()
}
