package codechallenge

import (
	"fmt"
	"strconv"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = false

// calculate() returns the result of the requested operation.
func calculateVariables(input1 string, input2 string, operation string) float64 {
	// Your code goes here.
	res := 0.0
	v1 := convertInputToValue(input1)
	v2 := convertInputToValue(input2)

	switch operation {
	case "+":
		res = addValues(v1, v2)
	case "-":
		res = subtractValues(v1, v2)
	case "*":
		res = multiplyValues(v1, v2)
	case "/":
		res = divideValues(v1, v2)
	default:
		fmt.Println("invalid operation")
	}
	return res
}

func convertInputToValue(input string) float64 {
	val, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		msg := fmt.Sprintf("%v must be a valid number", input)
		panic(msg)
	}
	return val
}

func addValues(value1, value2 float64) float64 {
	return value1 + value2
}

func subtractValues(value1, value2 float64) float64 {
	return value1 - value2
}

func multiplyValues(value1, value2 float64) float64 {
	return value1 * value2
}

func divideValues(value1, value2 float64) float64 {
	if value2 == 0 {
		msg := fmt.Sprintln("cannot be divided by zero")
		panic(msg)
	}
	return value1 / value2
}
