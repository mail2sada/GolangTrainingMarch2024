package main

import "fmt"

func RectanglePropery(a, b float64) (area float64, cir float64, str string) {
	area = a * b
	cir = 2 * (a + b)
	return
}

func main() {
	fmt.Println("Named return")

	a, c, s := RectanglePropery(1.0, 3.0)

	fmt.Println(a, c)
}

func
