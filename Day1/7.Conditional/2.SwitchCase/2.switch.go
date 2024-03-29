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
	case Sunday:
		fmt.Println("Sunday")
		fallthrough
	case Monday:
		fmt.Println("Monday")
	case Tuesday:
		fmt.Println("Tuesday")
	case Wednsday:
		fmt.Println("Wednsday")
	case Thursday:
		fmt.Println("Thursday")
	case Friday:
		fmt.Println("Friday")
	case Saturday:
		fmt.Println("Saturday")
		fallthrough
	default:
		fmt.Println("Invalid")
	}
}
