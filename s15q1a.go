package main

import "fmt"

func calculate(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}

func main() {

	resultSum, resultDiff := calculate(10, 5)

	fmt.Println("Sum:", resultSum)
	fmt.Println("Difference:", resultDiff)
}
