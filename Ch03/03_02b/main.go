package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")

	num := 42
	var p *int = &num

	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("Value of p:", *p)
	}

	val1 := 42.13
	p1 := &val1
	fmt.Println("Value 1 is", *p1)

	// modify the value through the pointer
	*p1 = *p1 / 31
	fmt.Println("New value is", val1)
}
