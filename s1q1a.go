package main

import "fmt"

func main() {
	var a, b int
	var opr string

	fmt.Println("Enter no`s a")
	fmt.Scanln(&a)
	fmt.Println("Enter operator")
	fmt.Scanln(&opr)
	fmt.Println("Enter no`s b")
	fmt.Scanln(&b)

	switch opr {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	default:
		fmt.Println("invalid opr!")
	}

}
