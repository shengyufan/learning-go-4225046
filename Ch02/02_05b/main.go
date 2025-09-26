package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	newSum := math.Round(sum*100) / 100
	fmt.Println("New sum:", newSum)

	radius := 5.5
	circumference := math.Pi * 2 * radius
	// 保留两位小数
	fmt.Printf("Circumference result is %.2f\n", circumference)
}
