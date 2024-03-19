package main

import "fmt"

func main() {

	fmt.Println("Working with structs...")

	type Employee struct {
		Eid        int
		Name, Dept string
	}

	var emp Employee = Employee{Eid: 1, Name: "Ajay..", Dept: "RCS"}

	fmt.Println(emp)
}
