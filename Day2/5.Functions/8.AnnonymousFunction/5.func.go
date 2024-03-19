package main

import "fmt"

func HOFunction() func() {

	a := func() {
		fmt.Println("This from function returned from HOF")
	}
	return a
}

func main() {

	fmt.Println("Higher order function")

	fn := HOFunction()

	fn()
	fn()
}
