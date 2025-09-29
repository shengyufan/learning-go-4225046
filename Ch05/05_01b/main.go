package main

import (
	"fmt"
)

func main() {
	fmt.Println("Functions")
	doSomething()
}

/*
commonly usecases:
1. func name starts with lowercase character is used in the same package
2. func name starts with uppercase one is public to the rest of the application
*/
func doSomething() {
	fmt.Println("Doing Something")
	v1 := 10
	v2 := 40
	sum := addValues(v1, v2)
	println(sum)
	sum = addAllValues(v1, v2, 50, 70, 2, 80)
	println(sum)
	sum, cnt, avg := addAllValuesWithMultipleReturns(v1, v2, 50, 70, 2, 80)
	println(sum, cnt, avg)
}

func addValues(val1 int, val2 int) int {
	return val1 + val2
}

func addAllValues(values ...int) int {
	sum := 0
	for _, v := range values {
		sum += v
	}
	return sum
}

// one function can return multiple values
func addAllValuesWithMultipleReturns(values ...int) (int, int, float64) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	count := len(values)
	avg := float64(sum) / float64(count)

	return sum, count, avg
}
