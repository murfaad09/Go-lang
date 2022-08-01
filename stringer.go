package main

import "fmt"

type Stringer struct {
	name string
	id   int
}

func (s Stringer) Stringer() string {

	fmt.Sprintln("name:")
	return s.Stringer()
}
