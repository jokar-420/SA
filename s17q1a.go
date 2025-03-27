package main

import "fmt"

func opr(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}
func main() {
	a, b := 20, 5
	sum, sub, mul, div := opr(a, b)
	fmt.Println("Sum:", sum, "\nSub", sub, "\nMultiplication", mul, "\nDivision", div)
}
