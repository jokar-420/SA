package main

import "fmt"

func Add_sub(a, b int) (int, int) {
	return a + b, a * b
}
func main() {
	a, b := 10, 20
	sum, mul := Add_sub(a, b)
	fmt.Println("Sum=", sum, "\nMultiplication=", mul)
}
