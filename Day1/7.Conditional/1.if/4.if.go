package main

import "fmt"

func main() {
	fmt.Println("Demo: if else if")

	const s, m, t, w, th, f, st = 1, 2, 3, 4, 5, 6, 7

	fmt.Print("Please enter weekday")

	var weekday int

	fmt.Scanf("%d", &weekday)

	if weekday != s {
		if weekday == m {

		}
	} else {
		fmt.Println("Sunday")
	}
}
