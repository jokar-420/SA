package main

import "fmt"

func main() {
	var n int
	t1, t2 := 0, 1
	fmt.Println("enter no.")
	fmt.Scanln(&n)
	fmt.Println(t1, "\n", t2)
	for i := 3; i <= n; i++ {
		t3 := t1 + t2
		fmt.Println(t3)
		t1 = t2
		t2 = t3

	}

}
