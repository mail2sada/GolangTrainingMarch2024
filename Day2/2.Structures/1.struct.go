package main

import "fmt"

/*
type <Name Struct> struct {
	f1 type
	f2 type
	f3 type
}
*/

type Student struct {
	Id    int
	Name  string
	Class string
	Dept  string
}

func main() {

	var std Student

	fmt.Println(std)

	std = Student{Id: 0, Name: "Anil kumar", Class: "10thB", Dept: "Maths"}

	fmt.Println(std)

	fmt.Println(std.Id, std.Name, std.Class)

	std.Id = 100

	fmt.Println(std)

}
