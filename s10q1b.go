package main

import "fmt"

// Function to generate Fibonacci series and send it to the channel
func fibonacci(n int, ch chan int) {
	a, b := 0, 1
	for i := 0; i < n; i++ {
		ch <- a
		a, b = b, a+b
	}
	close(ch) // Close the channel after sending all Fibonacci numbers
}

func main() {
	// Set the number of Fibonacci terms
	n := 10

	// Create a channel
	ch := make(chan int)

	// Start a goroutine to generate Fibonacci numbers
	go fibonacci(n, ch)

	// Print Fibonacci numbers received from the channel
	for num := range ch {
		fmt.Println(num)
	}
}
