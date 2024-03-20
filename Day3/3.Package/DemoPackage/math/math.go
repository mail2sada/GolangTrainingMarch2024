package math

import (
	"first/geo"
	"fmt"
)

const pI = 3.14159

func init() {
	fmt.Println("Init from math")
}
func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Dic(a, b int) int {
	return a / b
}

func MethodFromGeo() {
	fmt.Println(geo.AsianContries())
}
