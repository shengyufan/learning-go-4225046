package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an array
	var colors_arr = [3]string{"Red", "Green", "Blue"}
	fmt.Println(colors_arr)

	// This is a slice
	// init method 1: no specify number/size
	var colors_slice = []string{"Red", "Green", "Blue"}
	fmt.Println(colors_slice)

	// init method 2: make function to allocate the memory
	var colors = make([]string, 0, 3)
	colors = append(colors, "red", "green", "blue")
	colors = append(colors, "white", "black")
	fmt.Println(colors)

	// remove/delete the item
	colors = remove(colors, 1)
	fmt.Println(colors)

	// sort the slice
	sort.Strings(colors)
	fmt.Println(colors)
}

func remove(slice []string, i int) []string {
	// ... means return anything else
	return append(slice[:i], slice[i+1:]...)
}
