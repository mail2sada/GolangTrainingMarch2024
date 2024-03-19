package main

import (
	"fmt"
	"reflect"
)

func Add(a, b int, elem ...int) int {

	sum := a + b
	fmt.Println(reflect.TypeOf(elem))
	for _, v := range elem {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println("Variadic functions")

	fmt.Println(Add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(Add(1, 2))

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	res := Add(1, 2, a...)
	fmt.Println(res)

	// fmt.Println(Add(1))

	// fmt.Println(Add())
}
