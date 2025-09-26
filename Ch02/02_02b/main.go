package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."

	// fmt.Println("Hello from Go!")
	fmt.Println(str1, str2, str3)

	testNum := 42
	// print the integer number
	fmt.Printf("Test Number is: %v\n", testNum)
	// print the data type
	fmt.Printf("Test Data Type is: %T\n", testNum)
}
