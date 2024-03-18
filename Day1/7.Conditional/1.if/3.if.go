package main

import "fmt"

func main() {
	fmt.Println("Demo: if else if")

	const s, m, t, w, th, f, st = 1, 2, 3, 4, 5, 6, 7

	fmt.Print("Please enter weekday")

	var weekday int

	fmt.Scanf("%d", &weekday)

	if weekday == s {
		fmt.Println("Sun")
	} else if weekday == m {
		fmt.Println("Mon")

	} else if weekday == t {
		fmt.Println("tue")

	} else if weekday == w {
		fmt.Println("Wed")

	} else if weekday == th {
		fmt.Println("Thu")
	} else if weekday == f {
		fmt.Println("Fri")

	} else if weekday == st {
		fmt.Println("Sat")

	} else {
		fmt.Println("Invalid input")
	}
}
