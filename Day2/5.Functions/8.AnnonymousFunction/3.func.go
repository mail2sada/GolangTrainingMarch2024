package main

import "fmt"

var (
	a = func(msg string) {
		fmt.Println(msg)
	}
)

func main() {

	fmt.Println("Exploring annonymous function")
	fn := func(msg string) {
		fmt.Println(msg)
	}

	fn("Hi")
	fn("Hello")
	a("How are you?")
	fmt.Println("Exiting main")
}
