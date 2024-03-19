package main

import "fmt"

func main() {
	fmt.Println("Defer...")

	defer fmt.Println("We are calling this function defered")

	fmt.Println("After defer")
}
