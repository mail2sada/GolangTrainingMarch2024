package main

import "fmt"

/*
	const <name> <type> = <values>
*/

func main() {
	fmt.Println("Demo: Constants")

	const s, m, t, w, th, f, st = 1, 2, 3, 4, 5, 6, 7

	const a, b int = 10, 20

	fmt.Println("Enter a week number")

	var week = 7

	fmt.Println("week day is", week, week == s, week == st)

}
