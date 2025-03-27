package main

import "fmt"

type emp struct {
	eno     int
	ename   string
	esalary float64
}

func main() {
	var n int
	fmt.Println("Enter number of employees:")
	fmt.Scanln(&n)
	e := make([]emp, n)

	for i := 0; i < n; i++ {
		fmt.Println("Enter details of employee no.", i+1)
		fmt.Print("Employee No: ")
		fmt.Scanln(&e[i].eno)
		fmt.Print("Employee Name: ")
		fmt.Scanln(&e[i].ename)
		fmt.Print("Employee Salary: ")
		fmt.Scanln(&e[i].esalary)
	}

	// Initialize min to the first employee's salary
	min := e[0].esalary
	for j := 1; j < n; j++ {
		if min > e[j].esalary {
			min = e[j].esalary
		}
	}

	// Print employees with the minimum salary
	fmt.Println("\nEmployees with minimum salary:")
	for _, emp := range e {
		if emp.esalary == min {
			fmt.Printf("Employee No: %d, Name: %s, Salary: %.2f\n", emp.eno, emp.ename, emp.esalary)
		}
	}
}
