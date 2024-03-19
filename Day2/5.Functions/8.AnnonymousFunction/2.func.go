package main

import "fmt"

func main() {
	fmt.Println("Demo: Annonymous function")

	func(msg string) {
		fmt.Println("Printing msg", msg)
	}("Hello")

	fmt.Println("Exiting main")
}
