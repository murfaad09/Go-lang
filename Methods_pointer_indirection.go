package main

type car struct {
	reading int
	milage  int
}

func (c car) show(e int) int {
	c.reading = c.reading + 2289
	c.milage = c.milage / e
	total := (c.reading + c.milage)
	return total
}
func main() {
	p := &car
	p.reading
	show(20)

}
