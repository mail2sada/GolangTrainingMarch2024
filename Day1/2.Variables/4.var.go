package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Demo: Variables")

	var integer, stringsss = 100, "Hello"

	fmt.Println(integer, stringsss)

	var a, b, c, d = 1, 10.22, true, "hello 111"

	fmt.Println(a, b, c, d)

	fmt.Println(reflect.TypeOf(a), reflect.TypeOf(b), reflect.TypeOf(c), reflect.TypeOf(d))

}
