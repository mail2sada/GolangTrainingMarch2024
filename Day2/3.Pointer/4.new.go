package main

import "fmt"

type Student struct {
	RollNo int
	Name   string
	Class  string
}

func main() {
	fmt.Println("New function")

	v := Student{RollNo: 1, Name: "abc", Class: "xyz"}
	fmt.Println(v)

	p := new(Student)
	p.RollNo = 1
	p.Class = "test"
	p.Name = "best"

	fmt.Println(p)
}
