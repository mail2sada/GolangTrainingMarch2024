package main

import "fmt"

x := 100

func main() {
	fmt.Println("Demo: Short delcaration operator")

	var a, b = 10, 20

	var c = a + b

	fmt.Println(c)

	res := a + b

	fmt.Println(res)

	add, sub, mul, div, rem := a+b, a-b, a*b, a/b, a%b

	fmt.Println(add, sub, mul, div, rem)
}
