package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	println(dog.SpeakThreeTimes())
}

type Dog struct {
	Breed string
	Sound string
}

func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

func (d Dog) SpeakThreeTimes() string {
	// Sprintf() return a variable
	return fmt.Sprintf("%v! %v! %v!", d.Sound, d.Sound, d.Sound)
}
