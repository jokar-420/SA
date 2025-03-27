package main

import "fmt"

// Define the student struct
type student struct {
	name  string
	age   int
	grade string
}

// Method to display student details
// The receiver is a pointer to the student struct
func (s *student) show() {
	fmt.Printf("Name: %s\n", s.name)
	fmt.Printf("Age: %d\n", s.age)
	fmt.Printf("Grade: %s\n", s.grade)
}

func main() {
	// Create a student instance
	s1 := student{
		name:  "John Doe",
		age:   20,
		grade: "A",
	}

	// Call the show method with a pointer to the student struct
	s1.show()
}
