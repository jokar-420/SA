package main

import (
	"fmt"
)

func main() {
	// Declare variables for length and width
	var length, width float64

	// Get user input for length and width
	fmt.Print("Enter the length of the rectangle: ")
	fmt.Scan(&length)
	fmt.Print("Enter the width of the rectangle: ")
	fmt.Scan(&width)

	// Call the Area function from the rectangle package
	area := rectangle.Area(length, width)

	// Print the area
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
