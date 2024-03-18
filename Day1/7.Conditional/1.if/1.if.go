package main

import "fmt"

func main() {
	fmt.Println("Demo:If")

	var a int
	fmt.Print("Please enter a number:")
	fmt.Scanf("%d", &a)

	if a < 0 {
		fmt.Println("Negative number")
	}

	if a >= 0 {
		fmt.Println("Positive number")
	}
}
