package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
}

func main() {
	fmt.Println("Demo: make")

	v := make([]Student, 10)

	fmt.Println(v)
}
