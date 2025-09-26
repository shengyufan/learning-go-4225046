package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read the keyboard input
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	// specify the end character
	str, _ := reader.ReadString('\n')
	fmt.Println("The value of input: ", str)

	// read the keyboard input (number)
	fmt.Print("Enter a number: ")
	// '=' to use the existing variable
	str, _ = reader.ReadString('\n')
	num, err := strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("The number is: %v\n", num)
	}

}
