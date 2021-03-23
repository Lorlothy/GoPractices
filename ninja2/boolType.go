package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)

	a := 10
	b := 55
	fmt.Println(a == b)
	fmt.Println(a >= b)
	fmt.Println(a <= b)
	fmt.Println(a != b)
}