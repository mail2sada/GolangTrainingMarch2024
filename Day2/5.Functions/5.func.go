package main

import "fmt"

const PI = 3.14159

func CircleProperty(radius float64) (float64, float64) {
	area := PI * radius * radius

	circ := 2 * PI * radius

	return area, circ
}

func main() {
	fmt.Println("Demo: Multiple return")

	a, c := CircleProperty(12.3)

	fmt.Println(a, c)

	_, c = CircleProperty(1.0)

	fmt.Println(c)

	a, _ = CircleProperty(1.0)
	fmt.Println(a)

}
