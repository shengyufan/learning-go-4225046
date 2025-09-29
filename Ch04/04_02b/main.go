package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var res string
	switch dayNumber {
	case 1:
		res = "Monday"
	case 2:
		res = "Tuesday"
	case 3:
		res = "Wednesday"
	case 4:
		res = "Thursday"
	case 5:
		res = "Friday"
	default:
		res = "Weekend"
	}

	fmt.Println(res)

	// can also specify the param if the param only be used in this code block
	/*
		switch x := -42; {
		case x < 0:
			res = "less than zero"
		case x == 0:
			res = "equal to zero"
		default:
			res = "great than zero"
		}
	*/
	x := -42
	switch {
	case x < 0:
		res = "less than zero"
		// similar to break statement if switch-break statement
		// fallthrough
	case x == 0:
		res = "equal to zero"
	default:
		res = "great than zero"
	}

	fmt.Println(res)
}
