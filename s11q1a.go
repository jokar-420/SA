package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter a number")
	fmt.Scanln(&n)

	if (n <= 99 && n >= 10) || (n <= -10 && n >= -99) {
		fmt.Println("yes ! given no. is two digit")
	} else {
		fmt.Println("no ! given no. is not two digit")

	}

}
