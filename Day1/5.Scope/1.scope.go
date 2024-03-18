package main

import "fmt"

var gb = 1000 // global variables

func main() {
	fmt.Println("Demo: Variable scope")
	var a = 100 // local variable

	fmt.Println(a)

	fmt.Println(gb)

	Test()

	gb = 20000

	Test()
}

func Test() {
	var b = 200 //local variable
	fmt.Println(b)
	fmt.Println(gb)
}
