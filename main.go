package main

import (
	"fmt"
	"rectangle" // Import the rectangle package
)

func main() {
	// Length and Width of the rectangle
	length := 5.0
	width := 3.0

	// Call the Area function from the rectangle package
	area := rectangle.Area(length, width)

	// Output the result
	fmt.Printf("Area of the rectangle: %.2f\n", area)
}
