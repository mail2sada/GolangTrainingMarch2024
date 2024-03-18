package main

import (
	"fmt"
	"reflect"
)

/*
	var <variable_name> <type> = <values>
*/

func main() {
	fmt.Println("Demo: Variable name...")

	var a int

	fmt.Println("Value of a is ", a)

	var c, d = 10, 20

	fmt.Println("Value of c and d", c, d)

	fmt.Printf("Type of c and d are %T %T\n", c, d)

	fmt.Println("Type of c and d are ", reflect.TypeOf(c), reflect.TypeOf(d))

}
