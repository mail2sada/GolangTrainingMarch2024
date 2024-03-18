package main

import "fmt"

/*
Global

Package global

Local to function

local to block
*/

var a = 1000 // global
func main() {

	var b = 10

	fmt.Println(b)
	{
		var c = 100
		fmt.Println(c)
		fmt.Println(b)

	}
	fmt.Println(b)

	fmt.Println(a)
}
