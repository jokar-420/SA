package main

import "fmt"

func mul(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println("n*", i, "=", n*i)
	}
}
func main() {
	n := 10
	mul(n)
}
