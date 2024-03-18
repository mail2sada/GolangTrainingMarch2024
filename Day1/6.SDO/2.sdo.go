package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: SDO")

	a, b := 1, "hi"

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b))
}
