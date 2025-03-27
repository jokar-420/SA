package main

import "fmt"

func main() {
	// Create a channel
	ch := make(chan int)

	// Goroutine to send data to the channel
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i // Send data to the channel
		}
		close(ch) // Close the channel after sending data
	}()

	// Use for range loop to receive data from the channel
	for value := range ch {
		fmt.Println(value) // Print the received value
	}

	fmt.Println("Channel closed and all data received.")
}
