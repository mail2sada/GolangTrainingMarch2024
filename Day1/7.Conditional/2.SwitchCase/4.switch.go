package main

import "fmt"

func main() {

	fmt.Println("Demo: Switch case")
	var digit int
	fmt.Print("Enter a value for digit")

	fmt.Scanf("%d", &digit)

	switch {
	// case digit >= -999 && digit <= -100:
	// 	fmt.Println("3 digits")
	// case digit >= -99:
	// 	fmt.Println("2 digits")
	case digit <= 9:
		fmt.Println("Single digit")
	case digit <= 99:
		fmt.Println("2 digits")
	case digit <= 999:
		fmt.Println("3 digits")
	default:
		fmt.Println("I dont have code for this")
	}
}
