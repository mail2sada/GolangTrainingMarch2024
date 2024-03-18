package main

import "fmt"

/*

for <initialis>;condition;increment {
	statements
}
*/

func main() {
	fmt.Println("Demo: for loop")

	for i := 0; i <= 10; i++ {
		fmt.Println("Value of i", i)
	}
}
