package main

import "fmt"

func palin(n int) {
	if n < 0 {
		return
	}
	t := n
	rev := 0
	for n > 0 {
		rev = rev*10 + (n % 10)
		n /= 10
	}
	if t == rev {
		fmt.Println("no is palin")
	} else {
		fmt.Println("no is not palin")
	}
}
func main() {
	var n int
	fmt.Println("Enter no")
	fmt.Scanln(&n)

	palin(n)
}
