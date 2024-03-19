package main

import "fmt"

type Std struct {
	Roll_no int
	Name    string
}

func InitStudent() *Std {
	fmt.Println("Inside InitStudent")
	return new(Std)
}

var (
	my_var *Std = InitStudent()
)

func init() {
	fmt.Println("This is init")
	my_var = InitStudent()
}

func main() {

	fmt.Println("Exploring init")
}
