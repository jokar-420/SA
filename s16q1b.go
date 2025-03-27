package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Loop to print numbers from 0 to 10
	for i := 0; i <= 10; i++ {
		// Print the current number
		fmt.Println(i)

		// Generate a random delay between 0 and 250 milliseconds
		delay := time.Duration(rand.Intn(251)) * time.Millisecond

		// Wait for the generated delay
		time.Sleep(delay)
	}
}
