package codechallenge

// Write your answer here, and then test your code.
// Your job is to implement the findLargest() method.

// Change these boolean values to control whether you see
// the expected answer and/or hints.
// const showExpectedResult = false
// const showHints = true

// slicesToObjects() returns a slice of Color objects
func slicesToObjects(colorNames []string, hexValues []int) []Color {
	// Your code goes here.
	// Create a slice of Color objects
	// res := make([]Color, 0)
	res := make([]Color, 0, len(colorNames))

	for i := 0; i < len(colorNames); i++ {
		color_obj := Color{colorNames[i], hexValues[i]}
		res = append(res, color_obj)
	}

	return res
}

type Color struct {
	Name string
	Hex  int
}
