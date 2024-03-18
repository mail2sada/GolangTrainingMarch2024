package main

import "fmt"

/*
+ addition
- subtraction
/ division
* multiplicaiton
% reminder
*/

func main() {

	fmt.Println("Airthmatic operators")

	var a, b = 100, 200

	var addition = a + b

	var sub = b - a

	var mul = a * b

	var div = b / a

	var rem = a % b

	fmt.Println(addition, sub, mul, div, rem)

}
