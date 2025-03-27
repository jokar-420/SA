package main

import "fmt"

func main() {
	oddSum, evenSum := 0, 0
	for i := 1; i <= 100; i++ {
		if i%2 == 0 {
			evenSum += i
		} else {
			oddSum += i
		}
	}
	fmt.Println("Sum of ODD numbers bet 1 to 100 is ", oddSum,
		"\nSum of Even numbers bet 1 to 100 is ", evenSum)
}
