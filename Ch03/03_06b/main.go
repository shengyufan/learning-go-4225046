package main

import (
	"fmt"
)

func main() {
	fmt.Println("Structs")

	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)

	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	// +v represents including the field name and the corresponding value at the same time
	fmt.Printf("%+v\n", poodle)

	// change the value
	poodle.Weight = 29
	fmt.Println(poodle)
}

// define the struct
type Dog struct {
	Breed  string
	Weight int
}
