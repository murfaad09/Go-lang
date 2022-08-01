package main

import "fmt"

type Student struct {
	name    string
	class   int
	subject string
	grade   string
}

func main() {
	var var1 Student
	var1.name = "Arslan"
	var1.class = 12
	var1.subject = "Maths"
	var1.grade = "A"

	print(var1)
}
func print(print1 Student) {
	fmt.Println("Name:", print1.name)
	fmt.Println("Class:", print1.class)
	fmt.Println("Subject:", print1.subject)
	fmt.Println("Grade:", print1.grade)
}
