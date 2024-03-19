package main

import "fmt"

func main() {
	fmt.Println("Exploring defer")

	defer fmt.Println("1st defer")

	defer fmt.Println("Second defer")

	fmt.Println("Exiting main")
}
