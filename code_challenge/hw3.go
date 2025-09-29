package codechallenge

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false;
// const showHints = false;

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	// Your code goes here.
	res := 0.0

	for _, item := range cart {
		res += item.Price * float64(item.Quantity)
	}

	return res
}
