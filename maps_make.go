package main

import "fmt"

func main() {
	var a = make(map[string]string)
	a["brand"] = "FOrd"
	a["model"] = "AUto"
	fmt.Printf("a:/t%v", a)
}
