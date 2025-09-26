package main

import "fmt"

func main() {
	fmt.Println("Math")

	i1, i2, i3 := 22, 33, 44
	fmt.Println("sum: ", i1+i2+i3)
	f1, f2, f3 := 22.2, 33.3, 44.4
	fmt.Println("sum: ", f1+f2+f3)

	// go 不支持隐式转换，因此进行数学计算时，需要保证类型相同
	total := float64(i1) + f2
	fmt.Println("sum: ", total)

}
