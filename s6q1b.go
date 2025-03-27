package main

import "fmt"

func main() {
	a := []int{9, 8, 9, 3, 2, 8}
	ca := make([]int, len(a))
	for i := 0; i < len(a); i++ {
		ca[i] = a[i]

	}
	fmt.Println("Copying array", ca)
}
