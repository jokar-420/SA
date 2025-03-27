package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}

	// Demonstrating append: Adding new elements to the slice
	a = append(a, 6, 7, 8)
	fmt.Println("After appending 6, 7, 8:", a)

	copySlice := make([]int, len(a))
	copy(copySlice, a)
	fmt.Println("Copied slice:", copySlice)

	indexToRemove := 2
	a = append(a[:indexToRemove], a[indexToRemove+1:]...)
	fmt.Println("After removing element at index 2:", a)
}
