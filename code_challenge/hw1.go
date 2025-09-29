// Write your answer here, and then test your code.
package codechallenge

import (
	"fmt"
	"strconv"
	"strings"
)

// Uncomment this import section to use required Go packages

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = true

// calculate() returns the the result of adding 2 numbers
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	// Your code goes here.

	// Convert the first string to a float64
	n1, err1 := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err1 != nil {
		fmt.Println(err1)
	}

	// Convert the second string to a float64
	n2, err2 := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err2 != nil {
		fmt.Println(err2)
	}

	// Calculate and return the result

	return n1 + n2
}
