package main

import "fmt"

/*

 && and

 || or

 ! not

 ^
*/

func main() {
	fmt.Println("Demo logical operators")

	var a, b, c, d = 10, 20, 30, 40

	var res = a != b && c < d

	fmt.Println(res) // true

	res = a != b && c > d

	fmt.Println(res) // false

	res = a != b || c > d

	fmt.Println(res)

	res = a != b && c > b || a < c
}
