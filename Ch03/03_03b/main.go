package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")

	// array size cannot be extended after initialization
	// array item cannot be sorted
	// init method 1
	var colors [3]string
	colors[0] = "red"
	colors[1] = "green"
	colors[2] = "blue"
	fmt.Println(colors)

	// init method 2
	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	fmt.Println("the size of the number array is", len(numbers))
}
