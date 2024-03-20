package main

import "fmt"

// type Student struct {
// 	Roll int
// 	Name string
// }

/*
	func (<Receiver>)MethodName(<Arguments>)ReturnValues {

	}
*/

func (std Student) Print() {
	fmt.Println("\t\t\t____________________________")
	fmt.Println("\t\t\tStudent Details are")
	fmt.Println("\t\t\t____________________________")

	fmt.Println("\t\t\tRollNo:", std.Roll, "Name:", std.Name)
	fmt.Println("\t\t\t____________________________")

}

func main() {
	fmt.Println("Exploring methods")

	var std = Student{Roll: 1, Name: "Anand R"}

	std.Print()
}
