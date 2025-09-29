package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	answer := 42
	var res string

	if answer < 0 {
		res = "less than zero"
	} else if answer == 0 {
		res = "equal to zero"
	} else {
		res = "greater than zero"
	}

	fmt.Println(res)
}
