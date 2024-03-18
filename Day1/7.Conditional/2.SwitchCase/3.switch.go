package main

import "fmt"

/*
switch <variable> {
case expr:

			statemets
	}
*/
const (
	Sunday int = iota
	Monday
	Tuesday
	Wednsday
	Thursday
	Friday
	Saturday
)

func main() {
	fmt.Println("Demo:Switch case fallthrough")

	fmt.Println("Please enter weekday")

	var wday int

	fmt.Scanf("%d", &wday)

	switch wday {
	case Saturday, Sunday:
		fmt.Println("Weekend")
	case Monday, Tuesday:
		fmt.Println("Beggining of the week")
	case Wednsday:
		fmt.Println("Mid of the week")
	case Thursday, Friday:
		fmt.Println("Approaching weekedn")
	default:
		fmt.Println("Invalid input")
	}
}
