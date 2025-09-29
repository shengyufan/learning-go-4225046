package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Files")
	fileName := "./fromString.txt"
	file, err := os.Create(fileName)
	checkError(err)

	// Defer the closing of the file. This will happen after the function returns.
	defer file.Close()
	
	length, err := io.WriteString(file, "Hello from Go!")
	fmt.Printf("Wrote a file with %v characters\n", length)
	readFile(fileName)
}

func checkError(err error) {
	if err != nil {
		// stop execution and output the error message
		panic(err)
	}
}

func readFile(fileName string) {
	data, err := os.ReadFile(fileName)
	checkError(err)
	println("Data:", string(data))
}
