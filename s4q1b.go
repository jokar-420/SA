package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{7, 8, 9, 5, 8}
	fmt.Println("Original Array")
	fmt.Println(a)
	sort.Ints(a)
	fmt.Println("sorting Array", a)

}
