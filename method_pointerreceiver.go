package main

import "fmt"

type class struct {
	name    string
	age     int
	subject string
}

func (c class) stu() int {

	return c.age * 5

}

func (c *class) stu1(x int) {
	c.age = c.age * x

}

func main() {
	c := class{"ahmed", 23, "maths"}
	c.stu1(5)
	fmt.Println("Age:", c.age)
}
