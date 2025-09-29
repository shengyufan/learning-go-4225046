package codechallenge

import (
	"encoding/json"
	"strings"
)

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = false

// type CartItem struct {
// 	Name  string
// 	Price float64
// 	Quantity int
// }

// getCartFromJson() returns a slice containing cartItem objects.
func getCartFromJson(jsonString string) []CartItem {
	var cart []CartItem
	// Your code goes here.
	// generate the decoder based on the JSON content
	decoder := json.NewDecoder(strings.NewReader(jsonString))
	_, err := decoder.Token()
	checkError(err)

	// iterate the content and convert it to the CartItem type
	var item CartItem
	for decoder.More() {
		err := decoder.Decode(&item)
		checkError(err)
		cart = append(cart, item)
	}

	return cart
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
