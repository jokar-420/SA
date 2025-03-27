package main

import "fmt"

func swap(n1, n2 *int) {
	t := *n1
	*n1 = *n2
	*n2 = t
	fmt.Println("After swapping a=", *n1, "b=", *n2)
}
func main() {
	a, b := 10, 20
	fmt.Println("Before swapping a=", a, "b=", b)
	swap(&a, &b)

}
