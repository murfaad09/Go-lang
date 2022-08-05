package main

import "fmt"

func slect(t, z chan string) {
	x, y := "X value", "Y value"
	for {
		select {
		case t <- x:
			fmt.Println("XX")
		case z <- y:
			fmt.Println("YY")
			return
		}
	}
}
func main() {
	t := make(chan string)
	z := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-t)

		}
		fmt.Println(<-z)

	}()
	slect(t, z)

}
