package main

import (
	"fmt"
	"reflect"
)

/*
 == equal to

 != not equal to

 < less than

 > greater than

 <= less than or equal

 >= greater than or equal
*/

func main() {

	fmt.Println("Demo: Comparision Operators")

	var a, b = 10, 20

	var res = a == b

	fmt.Println(res) // false

	res = a != b

	fmt.Println(res) // true

	res = a < b

	fmt.Println(res) // true

	res = a > b

	fmt.Println(res) // false

	res = a <= b

	fmt.Println(res) // true

	res = a >= b

	fmt.Println(res) // false

	fmt.Println("Type of res is", reflect.TypeOf(res))

}
