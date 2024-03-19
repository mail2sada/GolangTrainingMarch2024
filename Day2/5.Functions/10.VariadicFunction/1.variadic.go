package main

import "fmt"

func Add(elem ...int) int {

	sum := 0

	for _, v := range elem {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Variadic functions")

	fmt.Println(Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(Add(1, 2))
	fmt.Println(Add(1))
	fmt.Println(Add())
}
